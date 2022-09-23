package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin"
	"github.com/loebfly/ezgin/engine"
	"os"
)

func main() {
	err := ezgin.Start()
	if err != nil {
		return
	}

	// 业务...

	ezgin.Engine.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})

	ezgin.Engine.Routers(engine.Get, map[string]engine.GinHandlerFunc{
		"/hello2": TestHandler,
		"/hello3": TestHandler,
	})

	ezgin.Engine.GroupRoutes(engine.Get, "group", map[string]engine.GinHandlerFunc{
		"hello4": TestHandler,
		"hello5": TestHandler,
	})

	ezgin.Engine.FreeRoutes(map[engine.HttpMethod]map[string]engine.GinHandlerFunc{
		engine.Get: {
			"/hello6": TestHandler,
			"/hello7": TestHandler,
		},
		engine.Post: {
			"/hello8": TestHandler,
			"/hello9": TestHandler,
		},
	})

	ezgin.Engine.FreeGroupRoutes(map[string]map[engine.HttpMethod]map[string]engine.GinHandlerFunc{
		"group2": {
			engine.Get: {
				"/hello6": TestHandler,
				"/hello7": TestHandler,
			},
			engine.Post: {
				"/hello8": TestHandler,
				"/hello9": TestHandler,
			},
		},
	})

	// 服务异常退出时 优雅关闭服务
	ezgin.ShutdownWhenException(func(signal os.Signal) {

	}, func(ctx context.Context) {

	})
}

func TestHandler(ctx *gin.Context) engine.GinResult {
	return engine.GinResult{}
}
