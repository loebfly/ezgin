package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/internal/engine/middleware/trace"
)

const (
	MWTrace = trace.Enter
)

type enter int

const Enter = enter(0)

func InitObj(obj Yml) {
	config.initObj(obj)
	ctl.initEngine()
}

// GetOriGin 获取原生gin.Engine
func (enter) GetOriGin() *gin.Engine {
	return ctl.engine
}

// GetMWTraceCurHeaders 获取当前请求的所有header
func (enter) GetMWTraceCurHeaders() map[string]string {
	return MWTrace.GetCurHeader()
}

// GetMWTraceCurHeaderValueFor 获取当前请求的指定header值
func (enter) GetMWTraceCurHeaderValueFor(key string) string {
	return MWTrace.GetCurHeader()[key]
}

// CopyMWTracePreHeaderToCurRoutine 复制上一个请求的header到当前请求
func (enter) CopyMWTracePreHeaderToCurRoutine(preRoutineId string) {
	MWTrace.CopyPreHeaderToCurRoutine(preRoutineId)
}

// GetMWTraceCurRoutineId 获取当前请求的routineId
func (enter) GetMWTraceCurRoutineId() string {
	return MWTrace.GetCurRoutineId()
}

// GetMWTraceCurReqId 获取当前请求的reqId
func (enter) GetMWTraceCurReqId() string {
	return MWTrace.GetCurReqId()
}

// GetMWTraceCurClientIP 获取当前请求的客户端IP
func (enter) GetMWTraceCurClientIP() string {
	return MWTrace.GetCurClientIP()
}

// GetMWTraceCurUserAgent 获取当前请求的UserAgent
func (receiver enter) GetMWTraceCurUserAgent() string {
	return MWTrace.GetCurUserAgent()
}

// GetMWTraceCurXLang 获取当前请求的XLang
func (receiver enter) GetMWTraceCurXLang() string {
	return MWTrace.GetCurXLang()
}

func (enter) Any(relativePath string, handler engine.HandlerFunc) EZRouter {
	return ctl.Routers(engine.Any, map[string]engine.HandlerFunc{relativePath: handler})
}

func (enter) Get(relativePath string, handler engine.HandlerFunc) EZRouter {
	return ctl.Routers(engine.Get, map[string]engine.HandlerFunc{relativePath: handler})
}

func (enter) Post(relativePath string, handler engine.HandlerFunc) EZRouter {
	return ctl.Routers(engine.Post, map[string]engine.HandlerFunc{relativePath: handler})
}

func (enter) Delete(relativePath string, handler engine.HandlerFunc) EZRouter {
	return ctl.Routers(engine.Delete, map[string]engine.HandlerFunc{relativePath: handler})
}

func (enter) Patch(relativePath string, handler engine.HandlerFunc) EZRouter {
	return ctl.Routers(engine.Patch, map[string]engine.HandlerFunc{relativePath: handler})
}

func (enter) Put(relativePath string, handler engine.HandlerFunc) EZRouter {
	return ctl.Routers(engine.Put, map[string]engine.HandlerFunc{relativePath: handler})
}

func (enter) Head(relativePath string, handler engine.HandlerFunc) EZRouter {
	return ctl.Routers(engine.Head, map[string]engine.HandlerFunc{relativePath: handler})
}

func (enter) Options(relativePath string, handler engine.HandlerFunc) EZRouter {
	return ctl.Routers(engine.Options, map[string]engine.HandlerFunc{relativePath: handler})
}

func (enter) Use(middleware ...engine.MiddlewareFunc) EZRouter {
	return ctl.Use(middleware...)
}
func (enter) Group(relativePath string) EZRouter {
	return ctl.Group(relativePath)
}
func (enter) Routers(method engine.HttpMethod, pathHandler map[string]engine.HandlerFunc) EZRouter {
	return ctl.Routers(method, pathHandler)
}
func (enter) FreeRouters(methodPathHandlers map[engine.HttpMethod]map[string]engine.HandlerFunc) EZRouter {
	return ctl.FreeRouters(methodPathHandlers)
}
