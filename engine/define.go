package engine

import "github.com/gin-gonic/gin"

type HttpMethod string

const (
	ErrorCodeServiceUnavailable  = 1010
	ErrorCodeResUnmarshalFailure = 1011
)

const (
	Any     HttpMethod = "ANY"
	Get     HttpMethod = "GET"
	Head    HttpMethod = "HEAD"
	Post    HttpMethod = "POST"
	Put     HttpMethod = "PUT"
	Patch   HttpMethod = "PATCH"
	Delete  HttpMethod = "DELETE"
	Options HttpMethod = "OPTIONS"
)

// MiddlewareFunc 中间件函数
type MiddlewareFunc func(c *gin.Context)

// HandlerFunc 路由处理函数
type HandlerFunc func(c *gin.Context) Result[any]

// RecoveryFunc 恢复函数
type RecoveryFunc func(c *gin.Context, err any)

// EZRouter 路由接口
type EZRouter interface {
	Use(middleware ...MiddlewareFunc) EZRouter
	Group(relativePath string) EZRouter
	Routers(method HttpMethod, pathHandler map[string]HandlerFunc) EZRouter
	FreeRouters(methodPathHandlers map[HttpMethod]map[string]HandlerFunc) EZRouter
	Any(relativePath string, handler HandlerFunc) EZRouter
	Get(relativePath string, handler HandlerFunc) EZRouter
	Head(relativePath string, handler HandlerFunc) EZRouter
	Post(relativePath string, handler HandlerFunc) EZRouter
	Put(relativePath string, handler HandlerFunc) EZRouter
	Patch(relativePath string, handler HandlerFunc) EZRouter
	Delete(relativePath string, handler HandlerFunc) EZRouter
	Options(relativePath string, handler HandlerFunc) EZRouter
}
