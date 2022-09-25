package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/config"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/klogs"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

var (
	servers = make([]*http.Server, 0)
)

// StartServer 启动服务
func (receiver enter) StartServer(ymlPath string, ginEngine *gin.Engine) error {

	err := config.Enter.Init(ymlPath)
	if err != nil {
		return err
	}

	err = klogs.Init(ymlPath)
	if err != nil {
		klogs.Warn("日志模块初始化错误:{}", err.Error())
	}

	engine.Enter.SetOriEngine(ginEngine)

	ez := config.ObjConfig.EZGin

	if ez.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	if ez.App.Port > 0 {
		// HTTP 端口
		servers = append(servers, &http.Server{
			Addr:    strconv.Itoa(ez.App.Port),
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
			Addr:    strconv.Itoa(ez.App.PortSsl),
			Handler: engine.Enter.GetOriEngine(),
		})
		if ez.App.Cert != "" && ez.App.Key != "" {
			go func() {
				path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
				if listenErr := servers[1].ListenAndServeTLS(path+"/"+ez.App.Cert, path+"/"+ez.App.Key); listenErr != nil {
					klogs.Error("侦听HTTP端口{}失败:{}", ez.App.PortSsl, listenErr.Error())
				}
			}()
		} else {
			klogs.Error("HTTPS证书私钥文件未配置")
		}
	}

	if ez.Nacos.Server != "" {

		//nacos.Enter.InitObj()
	}

	klogs.C("APP").Debug("|-----------------------------------|")
	klogs.C("APP").Debug("|	   			{} {}				|", ez.App.Name, ez.App.Version)
	klogs.C("APP").Debug("|-----------------------------------|")
	if ez.App.Port > 0 {
		klogs.C("APP").Debug("|		HTTP端口: {}				|", ez.App.Port)
	}
	if ez.App.PortSsl > 0 {
		klogs.C("APP").Debug("|		HTTPS端口: {}				|", ez.App.PortSsl)
	}

	return nil
}

// ShutdownWhenExitSignal 服务异常退出时 优雅关闭服务
func (receiver enter) ShutdownWhenExitSignal(will func(os.Signal), did func(context.Context)) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-signalChan
	klogs.Error("收到退出信号:{}", sig.String())
	klogs.Error("关闭服务...")
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
