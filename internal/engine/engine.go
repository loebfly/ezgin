package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/internal/engine/middleware"
	"github.com/loebfly/ezgin/internal/engine/middleware/trace"
	"github.com/loebfly/ezgin/internal/logs"
	"net/http"
	"path"
	"strings"
)

var ctl = new(control)

type EZRouter interface {
	Use(middleware ...engine.MiddlewareFunc) EZRouter
	Group(relativePath string) EZRouter
	Routers(method engine.HttpMethod, pathHandler map[string]engine.HandlerFunc) EZRouter
	FreeRouters(methodPathHandlers map[engine.HttpMethod]map[string]engine.HandlerFunc) EZRouter
}

type control struct {
	engine  *gin.Engine
	routers map[string]engine.HandlerFunc
}

func (receiver *control) initEngine() {
	receiver.routers = make(map[string]engine.HandlerFunc)
	receiver.engine = config.Gin.Engine
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
						"[" + requestId + "]": 1,
					}
				}
				return nil
			})
		}
		if strings.Contains(config.Gin.Middleware, "logs") {
			receiver.engine.Use(middleware.Logs(config.Gin.LogChan))
		}
		if strings.Contains(config.Gin.Middleware, "xlang") {
			receiver.engine.Use(middleware.XLang)
		}
		if strings.Contains(config.Gin.Middleware, "recover") {
			if config.Gin.RecoveryFunc != nil {
				config.Gin.RecoveryFunc = func(c *gin.Context, err interface{}) {
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
			receiver.engine.Use(middleware.Recover(config.Gin.RecoveryFunc))
		}
	}
}

func (receiver *control) lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

func (receiver *control) routersHandler(ctx *gin.Context) {
	for relativePath, handler := range receiver.routers {
		if strings.Contains(relativePath, ":") {
			for _, p := range ctx.Params {
				relativePath = strings.Replace(relativePath, ":"+p.Key, p.Value, -1)
			}
		}
		if ctx.Request.URL.Path == relativePath {
			ctx.JSON(http.StatusOK, handler(ctx))
			return
		}
	}

	ctx.JSON(http.StatusNotFound, engine.Result{
		Status:  -1,
		Message: "404 not found",
	})
}

func (receiver *control) Get(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Get, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *control) Post(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Post, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *control) Delete(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Delete, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *control) Patch(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Patch, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *control) Put(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Put, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *control) Head(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Head, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *control) Options(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Options, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *control) Use(middleware ...engine.MiddlewareFunc) EZRouter {
	ginHandlers := make([]gin.HandlerFunc, 0, len(middleware))
	for _, m := range middleware {
		ginHandlers = append(ginHandlers, gin.HandlerFunc(m))
	}
	receiver.engine.Use(ginHandlers...)
	return receiver
}

func (receiver *control) Group(relativePath string) EZRouter {
	basePath := ""
	if !strings.HasPrefix(relativePath, "/") {
		basePath += "/"
	}
	basePath += relativePath
	if receiver.lastChar(basePath) != '/' {
		basePath += "/"
	}

	return &groupControl{
		control:     receiver,
		basePath:    basePath,
		groupEngine: receiver.engine.Group(relativePath),
	}
}

// Routers 批量生成路由
func (receiver *control) Routers(method engine.HttpMethod, pathHandler map[string]engine.HandlerFunc) EZRouter {
	for relativePath, handler := range pathHandler {
		key := ""
		if !strings.HasPrefix(relativePath, "/") {
			key += "/"
		}
		key += relativePath
		receiver.routers[key] = handler
		switch method {
		case engine.Get:
			receiver.engine.GET(relativePath, receiver.routersHandler)
		case engine.Post:
			receiver.engine.POST(relativePath, receiver.routersHandler)
		case engine.Delete:
			receiver.engine.DELETE(relativePath, receiver.routersHandler)
		case engine.Patch:
			receiver.engine.PATCH(relativePath, receiver.routersHandler)
		case engine.Put:
			receiver.engine.PUT(relativePath, receiver.routersHandler)
		default:
			receiver.engine.Any(relativePath, receiver.routersHandler)
		}
	}
	return receiver
}

// FreeRouters 批量生成自由路由 map[请求方法]map[接口地址]处理函数
func (receiver *control) FreeRouters(methodPathHandlers map[engine.HttpMethod]map[string]engine.HandlerFunc) EZRouter {
	for method, pathHandler := range methodPathHandlers {
		receiver.Routers(method, pathHandler)
	}
	return receiver
}

/*** groupControl ****/

type groupControl struct {
	control     *control
	basePath    string
	groupEngine *gin.RouterGroup
}

func (receiver *groupControl) Get(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Get, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *groupControl) Post(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Post, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *groupControl) Delete(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Delete, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *groupControl) Patch(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Patch, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *groupControl) Put(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Put, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *groupControl) Head(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Head, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *groupControl) Options(relativePath string, handler engine.HandlerFunc) EZRouter {
	return receiver.Routers(engine.Options, map[string]engine.HandlerFunc{relativePath: handler})
}

func (receiver *groupControl) Use(middleware ...engine.MiddlewareFunc) EZRouter {
	ginHandlers := make([]gin.HandlerFunc, 0, len(middleware))
	for _, m := range middleware {
		ginHandlers = append(ginHandlers, gin.HandlerFunc(m))
	}
	receiver.groupEngine.Use(ginHandlers...)
	return receiver
}

func (receiver *groupControl) Group(relativePath string) EZRouter {
	return &groupControl{
		control:     receiver.control,
		basePath:    receiver.joinPaths(relativePath),
		groupEngine: receiver.groupEngine.Group(relativePath),
	}
}

func (receiver *groupControl) Routers(method engine.HttpMethod, pathHandler map[string]engine.HandlerFunc) EZRouter {
	for relativePath, handler := range pathHandler {
		key := receiver.basePath + relativePath
		receiver.control.routers[key] = handler
		switch method {
		case engine.Get:
			receiver.groupEngine.GET(relativePath, receiver.control.routersHandler)
		case engine.Post:
			receiver.groupEngine.POST(relativePath, receiver.control.routersHandler)
		case engine.Delete:
			receiver.groupEngine.DELETE(relativePath, receiver.control.routersHandler)
		case engine.Patch:
			receiver.groupEngine.PATCH(relativePath, receiver.control.routersHandler)
		case engine.Put:
			receiver.groupEngine.PUT(relativePath, receiver.control.routersHandler)
		default:
			receiver.groupEngine.Any(relativePath, receiver.control.routersHandler)
		}
	}
	return receiver
}

func (receiver *groupControl) FreeRouters(methodPathHandlers map[engine.HttpMethod]map[string]engine.HandlerFunc) EZRouter {
	for method, pathHandler := range methodPathHandlers {
		receiver.Routers(method, pathHandler)
	}
	return receiver
}

func (receiver *groupControl) joinPaths(relativePath string) string {
	finalPath := ""
	if receiver.basePath == "" {
		relativePath += "/"
	}
	finalPath = path.Join(receiver.basePath, relativePath)
	if receiver.lastChar(finalPath) != '/' {
		finalPath += "/"
	}
	return finalPath
}

func (receiver *groupControl) lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}
