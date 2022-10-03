package ezgin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/internal/app"
	"github.com/loebfly/ezgin/internal/cache"
	"github.com/loebfly/ezgin/internal/config"
	"github.com/loebfly/ezgin/internal/dblite"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/logs"
	"github.com/loebfly/ezgin/internal/nacos"
	"os"
)

const (
	Config = config.Enter // 配置
	Nacos  = nacos.Enter  // nacos
	Engine = engine.Enter // gin引擎
	Cache  = cache.Enter  // 缓存
	Logs   = logs.Enter   // 日志
	DBLite = dblite.Enter // 数据库
)

// StartWithEngine 自定义启动服务
// @param ymlPath yml配置文件路径, 为空时默认为当前程序所在目录的同名yml文件
// @param engine gin引擎, 传nil则使用gin默认引擎
func StartWithEngine(ymlPath string, engine *gin.Engine) {
	app.StartWithEngine(ymlPath, engine)
}

// Start 默认的方式启动服务
func Start() {
	app.StartWithEngine("", nil)
}

// ShutdownWhenExitSignalWithCallBack 服务异常退出时 优雅关闭服务
func ShutdownWhenExitSignalWithCallBack(will func(os.Signal), did func(context.Context)) {
	app.ShutdownWhenExitSignalWithCallBack(will, did)
}

// ShutdownWhenExitSignal 服务异常退出时 优雅关闭服务
func ShutdownWhenExitSignal() {
	app.ShutdownWhenExitSignalWithCallBack(nil, nil)
}
