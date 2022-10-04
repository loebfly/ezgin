package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/internal/engine/middleware"
	"github.com/loebfly/ezgin/internal/engine/middleware/trace"
	"github.com/loebfly/ezgin/internal/logs"
	"net/http"
	"strings"
)

var ctl = new(control)

type control struct {
	engine *gin.Engine
}

const (
	ContentTypeFormUrlEncode = "application/x-www-form-urlencoded"
	ContentTypeFormMultipart = "multipart/form-data"
)

func (receiver control) initEngine() {
	receiver.engine = config.Gin.Engine
	gin.SetMode(config.Gin.Mode)

	if config.Gin.Middleware != "-" {
		if strings.Contains(config.Gin.Middleware, "cors") {
			receiver.Use(middleware.Cors)
		}
		if strings.Contains(config.Gin.Middleware, "trace") {
			receiver.Use(middleware.Trace)
			logs.Use(func(category, level string) map[string]int {
				requestId := trace.Enter.GetCurReqId()
				if requestId != "" {
					return map[string]int{
						requestId: 2,
					}
				}
				return nil
			})
		}
		if strings.Contains(config.Gin.Middleware, "logs") {
			receiver.Use(middleware.Logs)
		}
	}

}

func (receiver control) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.Use(middleware...)
}

func (receiver control) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.Handle(httpMethod, relativePath, handlers...)
}

func (receiver control) Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.Any(relativePath, handlers...)
}

func (receiver control) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.GET(relativePath, handlers...)
}

func (receiver control) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.POST(relativePath, handlers...)
}

func (receiver control) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.DELETE(relativePath, handlers...)
}

func (receiver control) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.PATCH(relativePath, handlers...)
}

func (receiver control) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.PUT(relativePath, handlers...)
}

func (receiver control) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.OPTIONS(relativePath, handlers...)
}

func (receiver control) HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.HEAD(relativePath, handlers...)
}

func (receiver control) StaticFile(relativePath, filepath string) gin.IRoutes {
	return receiver.engine.StaticFile(relativePath, filepath)
}

func (receiver control) StaticFileFS(relativePath, filepath string, fs http.FileSystem) gin.IRoutes {
	return receiver.engine.StaticFileFS(relativePath, filepath, fs)
}

func (receiver control) Static(relativePath, root string) gin.IRoutes {
	return receiver.engine.Static(relativePath, root)
}

func (receiver control) StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes {
	return receiver.engine.StaticFS(relativePath, fs)
}

// Routers 批量生成路由
func (receiver control) Routers(method engine.HttpMethod, routers map[string]engine.HandlerFunc) gin.IRoutes {

	for path, handler := range routers {
		switch method {
		case engine.Get:
			receiver.engine.GET(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case engine.Post:
			receiver.engine.POST(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case engine.Delete:
			receiver.engine.DELETE(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case engine.Patch:
			receiver.engine.PATCH(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case engine.Put:
			receiver.engine.PUT(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		default:
			receiver.engine.Any(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		}
	}

	return receiver.engine
}

// GroupRoutes 批量生成路由组
func (receiver control) GroupRoutes(method engine.HttpMethod, group string, routers map[string]engine.HandlerFunc) gin.IRoutes {

	groupRouter := receiver.engine.Group(group)

	for path, handler := range routers {
		switch method {
		case engine.Get:
			groupRouter.GET(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case engine.Post:
			groupRouter.POST(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case engine.Delete:
			groupRouter.DELETE(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case engine.Patch:
			groupRouter.PATCH(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case engine.Put:
			groupRouter.PUT(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		default:
			groupRouter.Any(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		}
	}

	return groupRouter
}

// FreeRoutes 批量生成自由路由 map[请求方法]map[接口地址]处理函数
func (receiver control) FreeRoutes(routers map[engine.HttpMethod]map[string]engine.HandlerFunc) gin.IRoutes {
	for method, router := range routers {
		receiver.Routers(method, router)
	}
	return receiver.engine
}

// FreeGroupRoutes 批量生成自由路由组 map[路由组]map[请求方法]map[接口地址]处理函数
func (receiver control) FreeGroupRoutes(routers map[string]map[engine.HttpMethod]map[string]engine.HandlerFunc) gin.IRoutes {
	for group, groupRouter := range routers {
		for method, router := range groupRouter {
			receiver.GroupRoutes(method, group, router)
		}
	}
	return receiver.engine
}
