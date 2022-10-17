package ezgin

import (
	appDefine "github.com/loebfly/ezgin/app"
	"github.com/loebfly/ezgin/internal/app"
	"github.com/loebfly/ezgin/internal/cache"
	"github.com/loebfly/ezgin/internal/call"
	"github.com/loebfly/ezgin/internal/config"
	"github.com/loebfly/ezgin/internal/dblite"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/i18n"
	"github.com/loebfly/ezgin/internal/logs"
	"github.com/loebfly/ezgin/internal/nacos"
)

const (
	Config = config.Enter // 配置模块
	Nacos  = nacos.Enter  // nacos模块
	Engine = engine.Enter // gin模块
	Cache  = cache.Enter  // 缓存模块
	Logs   = logs.Enter   // 日志模块
	DBLite = dblite.Enter // 数据库模块
	Call   = call.Enter   // 微服务调用模块
	I18n   = i18n.Enter   // 国际化模块
)

// Start 启动服务
// @param start 启动配置
func Start(start ...appDefine.Start) {
	app.Start(start...)
}

// ShutdownWhenExitSignal 服务异常退出时 优雅关闭服务
// shutdown 退出配置
func ShutdownWhenExitSignal(shutdown ...appDefine.Shutdown) {
	app.ShutdownWhenExitSignal(shutdown...)
}
