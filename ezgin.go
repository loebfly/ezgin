package ezgin

import (
	"github.com/gin-gonic/gin"
	appDefine "github.com/loebfly/ezgin/app"
	engineDefine "github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/internal/app"
	"github.com/loebfly/ezgin/internal/engine"
)

// Start 启动服务
// @param start 启动配置
/*
	示例:
import (
	"github.com/loebfly/ezgin/app"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
func main() {
	ezgin.Start(app.Start{
		GinCfg: app.GinCfg{
			RecoveryHandler: func(c *gin.Context, err any) {
				c.JSON(http.StatusOK, i18n.SystemError.ErrorRes())
			},
			NoRouteHandler: func(c *gin.Context) {
				c.JSON(http.StatusOK, i18n.UrlNotFound.ErrorRes())
			},
			SwaggerRelativePath: "/docs/*any",
			SwaggerHandler:      ginSwagger.WrapHandler(swaggerFiles.Handler),
		},
	})
}
*/
func Start(start ...appDefine.Start) {
	app.Start(start...)
}

// ShutdownWhenExitSignal 服务异常退出时 优雅关闭服务
// shutdown 退出配置
func ShutdownWhenExitSignal(shutdown ...appDefine.Shutdown) {
	app.ShutdownWhenExitSignal(shutdown...)
}

// GetOriGin 获取原生gin.Engine
func GetOriGin() *gin.Engine {
	return engine.Enter.GetOriGin()
}

// GetMWTraceCurHeaders 获取当前请求的所有header
func GetMWTraceCurHeaders() map[string]string {
	return engine.Enter.GetMWTraceCurHeaders()
}

// GetMWTraceCurHeaderValueFor 获取当前请求的指定header值
func GetMWTraceCurHeaderValueFor(key string) string {
	return engine.Enter.GetMWTraceCurHeaderValueFor(key)
}

// CopyMWTracePreHeaderToCurRoutine 复制上一个请求的header到当前请求
func CopyMWTracePreHeaderToCurRoutine(preRoutineId string) {
	engine.Enter.CopyMWTracePreHeaderToCurRoutine(preRoutineId)
}

// GetMWTraceCurRoutineId 获取当前请求的routineId
func GetMWTraceCurRoutineId() string {
	return engine.Enter.GetMWTraceCurRoutineId()
}

// GetMWTraceCurReqId 获取当前请求的reqId
func GetMWTraceCurReqId() string {
	return engine.Enter.GetMWTraceCurReqId()
}

// GetMWTraceCurClientIP 获取当前请求的客户端IP
func GetMWTraceCurClientIP() string {
	return engine.Enter.GetMWTraceCurClientIP()
}

// GetMWTraceCurUserAgent 获取当前请求的UserAgent
func GetMWTraceCurUserAgent() string {
	return engine.Enter.GetMWTraceCurUserAgent()
}

// GetMWTraceCurXLang 获取当前请求的XLang
func GetMWTraceCurXLang() string {
	return engine.Enter.GetMWTraceCurXLang()
}

func Any(relativePath string, handler engineDefine.HandlerFunc) engine.EZRouter {
	return engine.Enter.Any(relativePath, handler)
}

func Get(relativePath string, handler engineDefine.HandlerFunc) engine.EZRouter {
	return engine.Enter.Get(relativePath, handler)
}

func Post(relativePath string, handler engineDefine.HandlerFunc) engine.EZRouter {
	return engine.Enter.Post(relativePath, handler)
}

func Delete(relativePath string, handler engineDefine.HandlerFunc) engine.EZRouter {
	return engine.Enter.Delete(relativePath, handler)
}

func Patch(relativePath string, handler engineDefine.HandlerFunc) engine.EZRouter {
	return engine.Enter.Patch(relativePath, handler)
}

func Put(relativePath string, handler engineDefine.HandlerFunc) engine.EZRouter {
	return engine.Enter.Put(relativePath, handler)
}

func Head(relativePath string, handler engineDefine.HandlerFunc) engine.EZRouter {
	return engine.Enter.Head(relativePath, handler)
}

func Options(relativePath string, handler engineDefine.HandlerFunc) engine.EZRouter {
	return engine.Enter.Options(relativePath, handler)
}

func Use(middleware ...engineDefine.MiddlewareFunc) engine.EZRouter {
	return engine.Enter.Use(middleware...)
}
func Group(relativePath string) engine.EZRouter {
	return engine.Enter.Group(relativePath)
}
func Routers(method engineDefine.HttpMethod, pathHandler map[string]engineDefine.HandlerFunc) engine.EZRouter {
	return engine.Enter.Routers(method, pathHandler)
}
func FreeRouters(methodPathHandlers map[engineDefine.HttpMethod]map[string]engineDefine.HandlerFunc) engine.EZRouter {
	return engine.Enter.FreeRouters(methodPathHandlers)
}
