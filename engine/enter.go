package engine

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type enter int

const Enter = enter(0)

type HttpMethod string

const (
	Any     HttpMethod = "ANY"
	Get     HttpMethod = "GET"
	Head    HttpMethod = "HEAD"
	Post    HttpMethod = "POST"
	Put     HttpMethod = "PUT"
	Patch   HttpMethod = "PATCH" // RFC 5789
	Delete  HttpMethod = "DELETE"
	Connect HttpMethod = "CONNECT"
	Options HttpMethod = "OPTIONS"
	Trace   HttpMethod = "TRACE"
)

type GinHandlerFunc func(ctx *gin.Context) GinResult

type GinResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    struct {
		Count int `json:"count"`
		Index int `json:"index"`
		Size  int `json:"size"`
		Total int `json:"total"`
	} `json:"page"`
}

func (enter) InitObj(obj Yml) {
	config.initObj(obj)
	gin.SetMode(config.Gin.Mode)
}

func (enter) GetOriEngine() *gin.Engine {
	return ctl.engine
}

func (enter) GetFormParams(ctx *gin.Context) map[string]string {
	return ctl.GetFormParams(ctx)
}

func (enter) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return ctl.Use(middleware...)
}

func (enter) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return ctl.Handle(httpMethod, relativePath, handlers...)
}

func (enter) Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return ctl.Any(relativePath, handlers...)
}

func (enter) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return ctl.GET(relativePath, handlers...)
}

func (enter) POST(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return ctl.POST(relativePath, handlers...)
}

func (enter) DELETE(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return ctl.DELETE(relativePath, handlers...)
}

func (enter) PATCH(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return ctl.PATCH(relativePath, handlers...)
}

func (enter) PUT(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return ctl.PUT(relativePath, handlers...)
}

func (enter) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return ctl.OPTIONS(relativePath, handlers...)
}

func (enter) HEAD(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return ctl.HEAD(relativePath, handlers...)
}

func (enter) StaticFile(relativePath, filepath string) gin.IRoutes {
	return ctl.StaticFile(relativePath, filepath)
}

func (enter) StaticFileFS(relativePath, filepath string, fs http.FileSystem) gin.IRoutes {
	return ctl.StaticFileFS(relativePath, filepath, fs)
}

func (enter) Static(relativePath, root string) gin.IRoutes {
	return ctl.Static(relativePath, root)
}

func (enter) StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes {
	return ctl.StaticFS(relativePath, fs)
}

// Routers 批量生成路由
func (enter) Routers(method HttpMethod, routers map[string]GinHandlerFunc) gin.IRoutes {
	return ctl.Routers(method, routers)
}

// GroupRoutes 批量生成路由组
func (enter) GroupRoutes(method HttpMethod, group string, routers map[string]GinHandlerFunc) gin.IRoutes {
	return ctl.GroupRoutes(method, group, routers)
}

// FreeRoutes 批量生成自由路由 map[请求方法]map[接口地址]处理函数
func (enter) FreeRoutes(routers map[HttpMethod]map[string]GinHandlerFunc) gin.IRoutes {
	return ctl.FreeRoutes(routers)
}

// FreeGroupRoutes 批量生成自由路由组 map[路由组]map[请求方法]map[接口地址]处理函数
func (enter) FreeGroupRoutes(routers map[string]map[HttpMethod]map[string]GinHandlerFunc) gin.IRoutes {
	return ctl.FreeGroupRoutes(routers)
}
