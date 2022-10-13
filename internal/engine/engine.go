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
	engine  *gin.Engine
	routers map[string]engine.HandlerFunc
}

func (receiver *control) initEngine() {
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
			receiver.Use(middleware.Logs(config.Gin.LogChan))
		}
		if strings.Contains(config.Gin.Middleware, "xlang") {
			receiver.Use(middleware.XLang)
		}
		if strings.Contains(config.Gin.Middleware, "recover") {
			if config.Gin.RecoveryFunc != nil {
				config.Gin.RecoveryFunc = func(c *gin.Context, err interface{}) {
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
			receiver.Use(middleware.Recover(config.Gin.RecoveryFunc))
		}
	}
}

func (receiver *control) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.Use(middleware...)
}

func (receiver *control) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.Handle(httpMethod, relativePath, handlers...)
}

func (receiver *control) Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.Any(relativePath, handlers...)
}

func (receiver *control) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.GET(relativePath, handlers...)
}

func (receiver *control) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.POST(relativePath, handlers...)
}

func (receiver *control) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.DELETE(relativePath, handlers...)
}

func (receiver *control) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.PATCH(relativePath, handlers...)
}

func (receiver *control) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.PUT(relativePath, handlers...)
}

func (receiver *control) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.OPTIONS(relativePath, handlers...)
}

func (receiver *control) HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return receiver.engine.HEAD(relativePath, handlers...)
}

func (receiver *control) StaticFile(relativePath, filepath string) gin.IRoutes {
	return receiver.engine.StaticFile(relativePath, filepath)
}

func (receiver *control) StaticFileFS(relativePath, filepath string, fs http.FileSystem) gin.IRoutes {
	return receiver.engine.StaticFileFS(relativePath, filepath, fs)
}

func (receiver *control) Static(relativePath, root string) gin.IRoutes {
	return receiver.engine.Static(relativePath, root)
}

func (receiver *control) StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes {
	return receiver.engine.StaticFS(relativePath, fs)
}

func (receiver *control) routersHandler(ctx *gin.Context) {
	result := receiver.routers[ctx.Request.URL.Path](ctx)
	ctx.JSON(http.StatusOK, result)
}

// Routers 批量生成路由
func (receiver *control) Routers(method engine.HttpMethod, routers map[string]engine.HandlerFunc) gin.IRoutes {
	for path, handler := range routers {
		receiver.routers["/"+path+"/"] = handler
		switch method {
		case engine.Get:
			receiver.engine.GET(path, receiver.routersHandler)
		case engine.Post:
			receiver.engine.POST(path, receiver.routersHandler)
		case engine.Delete:
			receiver.engine.DELETE(path, receiver.routersHandler)
		case engine.Patch:
			receiver.engine.PATCH(path, receiver.routersHandler)
		case engine.Put:
			receiver.engine.PUT(path, receiver.routersHandler)
		default:
			receiver.engine.Any(path, receiver.routersHandler)
		}
	}

	return receiver.engine
}

// GroupRoutes 批量生成路由组
func (receiver *control) GroupRoutes(method engine.HttpMethod, group string, routers map[string]engine.HandlerFunc) gin.IRoutes {
	groupRouter := receiver.engine.Group(group)
	for path, handler := range routers {
		receiver.routers["/"+group+"/"+path+"/"] = handler
		switch method {
		case engine.Get:
			groupRouter.GET(path, receiver.routersHandler)
		case engine.Post:
			groupRouter.POST(path, receiver.routersHandler)
		case engine.Delete:
			groupRouter.DELETE(path, receiver.routersHandler)
		case engine.Patch:
			groupRouter.PATCH(path, receiver.routersHandler)
		case engine.Put:
			groupRouter.PUT(path, receiver.routersHandler)
		default:
			groupRouter.Any(path, receiver.routersHandler)
		}
	}

	return groupRouter
}

// FreeRoutes 批量生成自由路由 map[请求方法]map[接口地址]处理函数
func (receiver *control) FreeRoutes(routers map[engine.HttpMethod]map[string]engine.HandlerFunc) gin.IRoutes {
	for method, router := range routers {
		receiver.Routers(method, router)
	}
	return receiver.engine
}

// FreeGroupRoutes 批量生成自由路由组 map[路由组]map[请求方法]map[接口地址]处理函数
func (receiver *control) FreeGroupRoutes(routers map[string]map[engine.HttpMethod]map[string]engine.HandlerFunc) []gin.IRoutes {
	var groupIR = make([]gin.IRoutes, 0)
	for group, groupRouter := range routers {
		for method, router := range groupRouter {
			groupIR = append(groupIR, receiver.GroupRoutes(method, group, router))
		}
	}
	return groupIR
}

// NoRoute adds handlers for NoRoute. It returns a 404 code by default.
func (receiver *control) NoRoute(handlers ...gin.HandlerFunc) {
	receiver.engine.NoRoute(handlers...)
}
