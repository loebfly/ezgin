package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/internal/engine/middleware/trace"
	"github.com/loebfly/ezgin/internal/engine/middleware/xlang"
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

// GetCurReqId 获取当前请求id
func (enter) GetCurReqId() string {
	return trace.Enter.GetCurReqId()
}

// GetCurXLang 获取当前请求语言
func (enter) GetCurXLang() string {
	return xlang.Enter.GetCurXLang()
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
