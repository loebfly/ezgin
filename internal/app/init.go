package app

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/internal/config"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/logs"
	"github.com/loebfly/ezgin/internal/nacos"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

// initPath 初始化所有组件入口
func (receiver enter) initEZGin(ymlPath string, ginEngine *gin.Engine) {
	receiver.initConfig(ymlPath)
	receiver.initLogs()
	receiver.initEngine(ginEngine)
	receiver.initServer()
	receiver.initNacos()
}

func (receiver enter) initConfig(ymlPath string) {
	if ymlPath == "" {
		ymlPath = receiver.getYml()
	}
	config.InitPath(ymlPath)
}

// initLogs 初始化日志模块
func (receiver enter) initLogs() {
	out := config.EZGin().Logs.Out
	file := config.EZGin().Logs.File
	if file == "" {
		file = "/opt/logs/" + config.EZGin().App.Name
	}
	yml := logs.Yml{
		Out:  out,
		File: file,
	}
	logs.InitObj(yml)
}

// initEngine 初始化gin引擎
func (receiver enter) initEngine(ginEngine *gin.Engine) {
	yml := engine.Yml{
		Mode:       config.EZGin().Gin.Mode,
		Middleware: config.EZGin().Gin.Middleware,
		Engine:     ginEngine,
	}
	engine.InitObj(yml)
}

// initServer 初始化服务
func (receiver enter) initServer() {
	ez := config.EZGin()

	if ez.App.Port > 0 {
		// HTTP 端口
		servers = append(servers, &http.Server{
			Addr:    ":" + strconv.Itoa(ez.App.Port),
			Handler: engine.Enter.GetOriEngine(),
		})
		go func() {
			if listenErr := servers[0].ListenAndServe(); listenErr != nil {
				panic(fmt.Errorf("侦听HTTP端口%d失败%s", ez.App.Port, listenErr.Error()))
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
					panic(fmt.Errorf("侦听HTTPS端口%d失败%s", ez.App.PortSsl, listenErr.Error()))
				}
			}()
		} else {
			panic("HTTPS端口启动失败: 证书文件不存在")
		}
	}
}

// initNacos 初始化nacos
func (receiver enter) initNacos() {
	ez := config.EZGin()
	if ez.Nacos.Server != "" && ez.Nacos.Yml.Nacos != "" {
		nacosPrefix := ez.Nacos.Yml.Nacos
		if nacosPrefix != "" {
			nacosUrl := ez.GetNacosUrl(nacosPrefix)
			var yml nacos.Yml
			err := config.Enter.GetYmlObj(nacosUrl, &yml)
			if err != nil {
				panic(fmt.Errorf("nacos配置文件获取失败: %s", err.Error()))
			}
			yml.App = nacos.AppYml{
				Name:    ez.App.Name,
				Ip:      ez.App.Ip,
				Port:    ez.App.Port,
				PortSsl: ez.App.PortSsl,
				Debug:   ez.App.Debug,
			}
			nacos.InitObj(yml)
		}
	}
}
