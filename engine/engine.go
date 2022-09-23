package engine

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var oriEngine *gin.Engine

func (receiver enter) GetOriEngine() *gin.Engine {
	return oriEngine
}

func (receiver enter) SetOriEngine(engine *gin.Engine) {
	oriEngine = engine
}

func (receiver enter) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return oriEngine.Use(middleware...)
}

func (receiver enter) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return oriEngine.Handle(httpMethod, relativePath, handlers...)
}

func (receiver enter) Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return oriEngine.Any(relativePath, handlers...)
}

func (receiver enter) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return oriEngine.GET(relativePath, handlers...)
}

func (receiver enter) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return oriEngine.POST(relativePath, handlers...)
}

func (receiver enter) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return oriEngine.DELETE(relativePath, handlers...)
}

func (receiver enter) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return oriEngine.PATCH(relativePath, handlers...)
}

func (receiver enter) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return oriEngine.PUT(relativePath, handlers...)
}

func (receiver enter) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return oriEngine.OPTIONS(relativePath, handlers...)
}

func (receiver enter) HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return oriEngine.HEAD(relativePath, handlers...)
}

func (receiver enter) StaticFile(relativePath, filepath string) gin.IRoutes {
	return oriEngine.StaticFile(relativePath, filepath)
}

func (receiver enter) StaticFileFS(relativePath, filepath string, fs http.FileSystem) gin.IRoutes {
	return oriEngine.StaticFileFS(relativePath, filepath, fs)
}

func (receiver enter) Static(relativePath, root string) gin.IRoutes {
	return oriEngine.Static(relativePath, root)
}

func (receiver enter) StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes {
	return oriEngine.StaticFS(relativePath, fs)
}

// Routers 批量生成路由
func (receiver enter) Routers(method HttpMethod, routers map[string]GinHandlerFunc) gin.IRoutes {

	for path, handler := range routers {
		switch method {
		case Get:
			oriEngine.GET(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case Post:
			oriEngine.POST(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case Delete:
			oriEngine.DELETE(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case Patch:
			oriEngine.PATCH(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case Put:
			oriEngine.PUT(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		default:
			oriEngine.Any(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		}
	}

	return oriEngine
}

// GroupRoutes 批量生成路由组
func (receiver enter) GroupRoutes(method HttpMethod, group string, routers map[string]GinHandlerFunc) gin.IRoutes {

	groupRouter := oriEngine.Group(group)

	for path, handler := range routers {
		switch method {
		case Get:
			groupRouter.GET(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case Post:
			groupRouter.POST(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case Delete:
			groupRouter.DELETE(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case Patch:
			groupRouter.PATCH(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case Put:
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
func (receiver enter) FreeRoutes(routers map[HttpMethod]map[string]GinHandlerFunc) gin.IRoutes {
	for method, router := range routers {
		receiver.Routers(method, router)
	}
	return oriEngine
}

// FreeGroupRoutes 批量生成自由路由组 map[路由组]map[请求方法]map[接口地址]处理函数
func (receiver enter) FreeGroupRoutes(routers map[string]map[HttpMethod]map[string]GinHandlerFunc) gin.IRoutes {
	for group, groupRouter := range routers {
		for method, router := range groupRouter {
			receiver.GroupRoutes(method, group, router)
		}
	}
	return oriEngine
}
