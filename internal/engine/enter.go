package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/internal/engine/middleware/trace"
	"github.com/loebfly/ezgin/internal/engine/middleware/xlang"
)

const (
	MWTrace = trace.Enter
	MWXLang = xlang.Enter
)

type enter int

const Enter = enter(0)

func InitObj(obj Yml) {
	config.initObj(obj)
	ctl.initEngine()
}

// CopyPreAllMwDataToCurRoutine 复制前一个中间件的数据到当前协程
func CopyPreAllMwDataToCurRoutine(preRoutineId string) {
	MWTrace.CopyPreAllToCurRoutine(preRoutineId)
	MWXLang.CopyPreXLangToCurRoutine(preRoutineId)
}

// GetOriGin 获取原生gin.Engine
func (enter) GetOriGin() *gin.Engine {
	return ctl.engine
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
