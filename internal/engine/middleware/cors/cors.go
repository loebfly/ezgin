package cors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Middleware 跨域中间件
func (enter) Middleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	allowHeaders := ""
	// 允许所有请求头
	for k := range ctx.Request.Header {
		allowHeaders += k + " ,"
	}
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", allowHeaders)
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
}
