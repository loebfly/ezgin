package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/engine/middleware"
	"github.com/loebfly/ezgin/engine/middleware/trace"
	"github.com/loebfly/ezgin/logs"
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
	receiver.engine = config.Gin.engine
	gin.SetMode(config.Gin.Mode)
	if config.Gin.Middleware != "-" {
		if strings.Contains(config.Gin.Middleware, "cors") {
			receiver.engine.Use(middleware.Cors)
		}
		if strings.Contains(config.Gin.Middleware, "trace") {
			receiver.engine.Use(middleware.Trace)
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
			receiver.engine.Use(middleware.Logs)
		}
	}
}

func (receiver control) GetFormParams(ctx *gin.Context) map[string]string {
	params := make(map[string]string)
	cType := ctx.ContentType()
	if cType != ContentTypeFormUrlEncode &&
		cType != ContentTypeFormMultipart {
		return params
	}
	if ctx.Request == nil {
		return params
	}
	if ctx.Request.Method == "GET" {
		for k, v := range ctx.Request.URL.Query() {
			params[k] = v[0]
		}
		return params
	} else {
		err := ctx.Request.ParseForm()
		if err != nil {
			return params
		}
		for k, v := range ctx.Request.PostForm {
			params[k] = v[0]
		}
		for k, v := range ctx.Request.URL.Query() {
			params[k] = v[0]
		}
		return params
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
func (receiver control) Routers(method HttpMethod, routers map[string]GinHandlerFunc) gin.IRoutes {

	for path, handler := range routers {
		switch method {
		case Get:
			receiver.engine.GET(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case Post:
			receiver.engine.POST(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case Delete:
			receiver.engine.DELETE(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case Patch:
			receiver.engine.PATCH(path, func(ctx *gin.Context) {
				result := handler(ctx)
				ctx.JSON(result.Code, result)
			})
		case Put:
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
func (receiver control) GroupRoutes(method HttpMethod, group string, routers map[string]GinHandlerFunc) gin.IRoutes {

	groupRouter := receiver.engine.Group(group)

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
func (receiver control) FreeRoutes(routers map[HttpMethod]map[string]GinHandlerFunc) gin.IRoutes {
	for method, router := range routers {
		receiver.Routers(method, router)
	}
	return receiver.engine
}

// FreeGroupRoutes 批量生成自由路由组 map[路由组]map[请求方法]map[接口地址]处理函数
func (receiver control) FreeGroupRoutes(routers map[string]map[HttpMethod]map[string]GinHandlerFunc) gin.IRoutes {
	for group, groupRouter := range routers {
		for method, router := range groupRouter {
			receiver.GroupRoutes(method, group, router)
		}
	}
	return receiver.engine
}
