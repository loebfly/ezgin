package app

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/config"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/nacos"
	"github.com/loebfly/klogs"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	servers = make([]*http.Server, 0)
)

// getLocalYml 获取yml配置文件路径
func (receiver enter) getYml() string {
	var fileName string
	flag.StringVar(&fileName, "f", os.Args[0]+".yml", "yml配置文件名")
	flag.Parse()
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if strings.Contains(fileName, "/") || strings.Contains(fileName, "\\") {
		return fileName
	}
	return path + "/" + fileName
}

// StartServer 启动服务
func (receiver enter) StartServer(ginEngine *gin.Engine) error {
	ymlPath := receiver.getYml()
	err := config.Enter.Init(ymlPath)
	if err != nil {
		return err
	}

	err = klogs.Init(ymlPath)
	if err != nil {
		klogs.Warn("日志模块初始化错误:{}", err.Error())
	}

	engine.Enter.SetOriEngine(ginEngine)

	ez := config.YmlObj.EZGin

	if ez.Gin.Mode != "" {
		gin.SetMode(ez.Gin.Mode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	if ez.App.Port > 0 {
		// HTTP 端口
		servers = append(servers, &http.Server{
			Addr:    ":" + strconv.Itoa(ez.App.Port),
			Handler: engine.Enter.GetOriEngine(),
		})
		go func() {
			if listenErr := servers[0].ListenAndServe(); listenErr != nil {
				klogs.Error("侦听HTTP端口{}失败:{}", ez.App.Port, listenErr.Error())
			}
		}()
	}
	if ez.App.PortSsl > 0 {
		// HTTPS 端口
		servers = append(servers, &http.Server{
			Addr:    ":" + strconv.Itoa(ez.App.PortSsl),
			Handler: engine.Enter.GetOriEngine(),
		})
		if ez.App.Cert != "" && ez.App.Key != "" {
			go func() {
				path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
				if listenErr := servers[1].ListenAndServeTLS(path+"/"+ez.App.Cert, path+"/"+ez.App.Key); listenErr != nil {
					klogs.Error("侦听HTTPS端口{}失败:{}", ez.App.PortSsl, listenErr.Error())
				}
			}()
		} else {
			klogs.Error("HTTPS证书私钥文件未配置")
		}
	}

	if ez.Nacos.Server != "" {
		nacosPrefix := config.YmlObj.EZGin.Nacos.Yml.Nacos
		if nacosPrefix != "" {
			nacosUrl := config.YmlObj.GetNacosUrl(config.YmlObj.EZGin.Nacos.Yml.Nacos)
			var yml nacos.Yml
			err = config.Enter.GetYmlObj(nacosUrl, &yml)
			if err != nil {
				klogs.Error("获取nacos配置失败:{}", err.Error())
				return err
			}

			err = nacos.Enter.InitObj(yml)
			if err != nil {
				klogs.Error("nacos初始化失败:{}", err.Error())
				return err
			}
		}
	}

	klogs.C("APP").Debug("|-----------------------------------|")
	klogs.C("APP").Debug("|	   			{} {}				|", ez.App.Name, ez.App.Version)
	klogs.C("APP").Debug("|-----------------------------------|")
	if ez.App.Port > 0 {
		klogs.C("APP").Debug("|	启动成功!		HTTP端口: {}				|", ez.App.Port)
	}
	if ez.App.PortSsl > 0 {
		klogs.C("APP").Debug("|	启动成功!		HTTPS端口: {}			|", ez.App.PortSsl)
	}
	klogs.C("APP").Debug("|-----------------------------------|")

	return nil
}

// ShutdownWhenExitSignal 服务异常退出时 优雅关闭服务
func (receiver enter) ShutdownWhenExitSignal(will func(os.Signal), did func(context.Context)) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-signalChan
	klogs.Error("收到退出信号:{}", sig.String())
	klogs.Error("关闭服务...")

	nacos.Enter.Unregister()

	will(sig)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, server := range servers {
		if server != nil {
			if err := server.Shutdown(ctx); err != nil {
				klogs.Error("关闭服务失败:{}", err.Error())
				return
			}
		}
	}

	klogs.Error("服务已关闭")
	did(ctx)
}
