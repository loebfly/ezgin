package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/internal/engine/middleware/cors"
	"github.com/loebfly/ezgin/internal/engine/middleware/reqlogs"
	"github.com/loebfly/ezgin/internal/engine/middleware/trace"
)

// Cors 跨域中间件
func Cors(ctx *gin.Context) {
	cors.Enter.Middleware(ctx)
}

func Trace(ctx *gin.Context) {
	trace.Enter.Middleware(ctx)
}

func Logs(logChan chan reqlogs.ReqCtx) func(ctx *gin.Context) {
	reqlogs.Enter.SetLogChan(logChan)
	return func(ctx *gin.Context) {
		reqlogs.Enter.Middleware(ctx)
	}
}