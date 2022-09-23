package ezgin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/app"
	"github.com/loebfly/ezgin/config"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/nacos"
	"os"
	"path/filepath"
)

const (
	App    = app.Enter
	Config = config.Enter
	Nacos  = nacos.Enter
	Engine = engine.Enter
)

// GetYmlPath 获取yml配置文件路径
func GetYmlPath(fileName string) string {
	if fileName != "" {
		path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		return path + fileName + ".yml"
	} else {
		return os.Args[0] + ".yml"
	}
}

// StartWithCustom 自定义启动服务
// @param ymlPath yml配置文件路径, 为空时默认为当前程序所在目录的同名yml文件
// @param engine gin引擎, 传nil则使用gin默认引擎
func StartWithCustom(ymlPath string, engine *gin.Engine) error {
	if ymlPath == "" {
		ymlPath = GetYmlPath("")
	}
	return App.StartServer(ymlPath, engine)
}

// Start 默认的方式启动服务
func Start() error {
	return StartWithCustom("", nil)
}

// ShutdownWhenExitSignalWithCallBack 服务异常退出时 优雅关闭服务
func ShutdownWhenExitSignalWithCallBack(will func(os.Signal), did func(context.Context)) {
	App.ShutdownWhenExitSignal(will, did)
}

// ShutdownWhenExitSignal 服务异常退出时 优雅关闭服务
func ShutdownWhenExitSignal() {
	ShutdownWhenExitSignalWithCallBack(nil, nil)
}
