package ezgin

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/internal/app"
	"github.com/loebfly/ezgin/internal/cache"
	"github.com/loebfly/ezgin/internal/call"
	"github.com/loebfly/ezgin/internal/config"
	"github.com/loebfly/ezgin/internal/dblite"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/i18n"
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
	Call   = call.Enter   // 微服务调用
	I18n   = i18n.Enter   // 国际化
)

// Start 完整参数启动服务
// @param ymlPath yml配置文件路径, 为空时默认为当前程序所在目录的同名yml文件
// @param engine gin引擎, 传nil则使用gin默认引擎
// @param recoveryFunc 异常回调, 传nil则使用gin默认回调
func Start(ymlPath string, engine *gin.Engine, recoveryFunc gin.RecoveryFunc) {
	app.StartWithEngine(ymlPath, engine, recoveryFunc)
}

func StartWithDefault() {
	app.StartWithEngine("", nil, nil)
}

// StartWithRecover 启动服务, 并捕获异常
// @param recoveryFunc 异常回调, 传nil则使用gin默认回调
func StartWithRecover(recoveryFunc gin.RecoveryFunc) {
	app.StartWithEngine("", nil, recoveryFunc)
}

// StartWithEngine 启动服务
// @param engine gin引擎, 传nil则使用gin默认引擎
func StartWithEngine(engine *gin.Engine) {
	app.StartWithEngine("", engine, nil)
}

// StartWithYml 启动服务
// @param ymlPath yml配置文件路径, 为空时默认为当前程序所在目录的同名yml文件
func StartWithYml(ymlPath string) {
	app.StartWithEngine(ymlPath, nil, nil)
}

// ShutdownWhenExitSignal 服务异常退出时 优雅关闭服务
func ShutdownWhenExitSignal(will func(os.Signal), did func(context.Context)) {
	app.ShutdownWhenExitSignalWithCallBack(will, did)
}

// ShutdownWhenExitSignalWithDefault 服务异常退出时 优雅关闭服务
func ShutdownWhenExitSignalWithDefault() {
	app.ShutdownWhenExitSignalWithCallBack(nil, nil)
}

// ShutdownWhenExitSignalWithWill 服务异常退出时 优雅关闭服务
func ShutdownWhenExitSignalWithWill(will func(os.Signal)) {
	app.ShutdownWhenExitSignalWithCallBack(will, nil)
}

// ShutdownWhenExitSignalWithDid 服务异常退出时 优雅关闭服务
func ShutdownWhenExitSignalWithDid(did func(context.Context)) {
	app.ShutdownWhenExitSignalWithCallBack(nil, did)
}
