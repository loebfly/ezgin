package ezgin

import (
	"github.com/gin-gonic/gin"
	appDefine "github.com/loebfly/ezgin/app"
	engineDefine "github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/internal/app"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/engine/middleware/trace"
)

const (
	MWTrace = trace.Enter
)

// Start 启动服务
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

func Any(relativePath string, handler engineDefine.HandlerFunc) engineDefine.EZRouter {
	return engine.Enter.Any(relativePath, handler)
}

func Get(relativePath string, handler engineDefine.HandlerFunc) engineDefine.EZRouter {
	return engine.Enter.Get(relativePath, handler)
}

func Post(relativePath string, handler engineDefine.HandlerFunc) engineDefine.EZRouter {
	return engine.Enter.Post(relativePath, handler)
}

func Delete(relativePath string, handler engineDefine.HandlerFunc) engineDefine.EZRouter {
	return engine.Enter.Delete(relativePath, handler)
}

func Patch(relativePath string, handler engineDefine.HandlerFunc) engineDefine.EZRouter {
	return engine.Enter.Patch(relativePath, handler)
}

func Put(relativePath string, handler engineDefine.HandlerFunc) engineDefine.EZRouter {
	return engine.Enter.Put(relativePath, handler)
}

func Head(relativePath string, handler engineDefine.HandlerFunc) engineDefine.EZRouter {
	return engine.Enter.Head(relativePath, handler)
}

func Options(relativePath string, handler engineDefine.HandlerFunc) engineDefine.EZRouter {
	return engine.Enter.Options(relativePath, handler)
}

func Use(middleware ...engineDefine.MiddlewareFunc) engineDefine.EZRouter {
	return engine.Enter.Use(middleware...)
}
func Group(relativePath string) engineDefine.EZRouter {
	return engine.Enter.Group(relativePath)
}
func Routers(method engineDefine.HttpMethod, pathHandler map[string]engineDefine.HandlerFunc) engineDefine.EZRouter {
	return engine.Enter.Routers(method, pathHandler)
}
func FreeRouters(methodPathHandlers map[engineDefine.HttpMethod]map[string]engineDefine.HandlerFunc) engineDefine.EZRouter {
	return engine.Enter.FreeRouters(methodPathHandlers)
}
