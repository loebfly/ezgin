package app

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/config"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/logs"
	"github.com/loebfly/ezgin/nacos"
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
	if ymlPath == "" {
		ymlPath = receiver.getYml()
	}
	config.Enter.InitPath(ymlPath)
	receiver.initLogs()
	receiver.initEngine(ginEngine)
	receiver.initServer()
	receiver.initNacos()
}

// initLogs 初始化日志模块
func (receiver enter) initLogs() {

	out := config.YmlObj.EZGin.Logs.Out
	file := config.YmlObj.EZGin.Logs.File
	if file == "" {
		file = "/opt/logs/" + config.YmlObj.EZGin.App.Name
	}
	yml := logs.Yml{
		Out:  out,
		File: file,
	}
	err := logs.Enter.InitObj(yml)
	if err != nil {
		logs.Enter.Warn("日志模块初始化错误:{}", err.Error())
	}
}

// initEngine 初始化gin引擎
func (receiver enter) initEngine(ginEngine *gin.Engine) {
	engine.Enter.SetOriEngine(ginEngine)

	ez := config.YmlObj.EZGin

	if ez.Gin.Mode != "" {
		gin.SetMode(ez.Gin.Mode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

// initServer 初始化服务
func (receiver enter) initServer() {
	ez := config.YmlObj.EZGin

	if ez.App.Port > 0 {
		// HTTP 端口
		servers = append(servers, &http.Server{
			Addr:    ":" + strconv.Itoa(ez.App.Port),
			Handler: engine.Enter.GetOriEngine(),
		})
		go func() {
			if listenErr := servers[0].ListenAndServe(); listenErr != nil {
				logs.Enter.Error("侦听HTTP端口{}失败:{}", ez.App.Port, listenErr.Error())
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
					logs.Enter.Error("侦听HTTPS端口{}失败:{}", ez.App.PortSsl, listenErr.Error())
				}
			}()
		} else {
			logs.Enter.Error("HTTPS证书私钥文件未配置")
		}
	}
}

// initNacos 初始化nacos
func (receiver enter) initNacos() {
	ez := config.YmlObj.EZGin

	if ez.Nacos.Server != "" {
		nacosPrefix := config.YmlObj.EZGin.Nacos.Yml.Nacos
		if nacosPrefix != "" {
			nacosUrl := config.YmlObj.GetNacosUrl(config.YmlObj.EZGin.Nacos.Yml.Nacos)
			var yml nacos.Yml
			err := config.Enter.GetYmlObj(nacosUrl, &yml)
			if err != nil {
				panic(fmt.Errorf("nacos配置文件获取失败: %s", err.Error()))
			}

			err = nacos.Enter.InitObj(yml)
			if err != nil {
				panic(fmt.Errorf("nacos初始化失败: %s", err.Error()))
			}
		}
	}
}
