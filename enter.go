package ezgin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/app"
	"github.com/loebfly/ezgin/cache"
	"github.com/loebfly/ezgin/config"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/nacos"
	"os"
)

const (
	Config = config.Enter
	Nacos  = nacos.Enter
	Engine = engine.Enter
	Cache  = cache.Enter
)

// StartWithEngine 自定义启动服务
// @param ymlPath yml配置文件路径, 为空时默认为当前程序所在目录的同名yml文件
// @param engine gin引擎, 传nil则使用gin默认引擎
func StartWithEngine(ymlPath string, engine *gin.Engine) {
	app.Enter.Start(ymlPath, engine)
}

// Start 默认的方式启动服务
func Start() {
	app.Enter.Start("", nil)
}

// ShutdownWhenExitSignalWithCallBack 服务异常退出时 优雅关闭服务
func ShutdownWhenExitSignalWithCallBack(will func(os.Signal), did func(context.Context)) {
	app.Enter.ShutdownWhenExitSignal(will, did)
}

// ShutdownWhenExitSignal 服务异常退出时 优雅关闭服务
func ShutdownWhenExitSignal() {
	app.Enter.ShutdownWhenExitSignal(nil, nil)
}
