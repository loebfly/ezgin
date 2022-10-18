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
	// Config 配置模块
	/*
		说明: 用于获取配置信息
			示例: Config.GetString("app.name")
	*/
	Config = config.Enter
	// Nacos 注册中心模块
	/*
		说明: 用于获取注册中心信息
			示例:
				Nacos.GetClient() 获取nacos客户端
				Nacos.GetService("ezgin") 获取微服务服务器地址和端口
	*/
	Nacos = nacos.Enter
	// Engine gin引擎模块
	/*
		说明: 主要用于路由注册、中间件等
			示例:
				Engine.Use(middleware) 注册中间件
				Engine.Group("bucket").Routers(engine.Get, map[string]engine.HandlerFunc{
					"new": TestHandler,
				}) 注册路由
	*/
	Engine = engine.Enter
	// Cache 内存缓存模块
	/*
		说明: 用于内存缓存
			示例:
				Cache.Table("ezgin").Add("key", "value", 0) 设置缓存
				Cache.Table("ezgin").Get("key") 获取缓存
	*/
	Cache = cache.Enter
	// Logs 日志模块
	/*
		说明: 用于日志打印和输入
			示例:
				Logs.Debug("test") 打印debug日志
				Logs.CDebug("category", "test") 打印debug日志
	*/
	Logs = logs.Enter
	// DBLite 数据库模块
	/*
		说明: 提供Mysql、Redis、MongoDB数据库操作
			示例:
				DBLite.Mysql() 获取mysql数据库操作对象
				DBLite.Redis() 获取redis数据库操作对象
				DBLite.Mongo() 获取mongo数据库操作对象
	*/
	DBLite = dblite.Enter
	// Call 微服务调用模块
	/*
		说明: 用于调用其他微服务，支持form、json、restful三种方式
	*/
	Call = call.Enter
	// I18n 国际化模块
	/*
		说明: 用于获取国际化信息
			示例:
				I18n.String("messageId") 获取当前语言的国际化内容
	*/
	I18n = i18n.Enter
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
