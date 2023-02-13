package engine

import "github.com/gin-gonic/gin"

const (
	ContentTypeFormUrlEncode = "application/x-www-form-urlencoded"
	ContentTypeFormMultipart = "multipart/form-data"
)

func GetFormParams(ctx *gin.Context) map[string]string {
	params := make(map[string]string)
	cType := ctx.ContentType()
	// bugfix: ctx.ContentType() 可能为空，导致无法获取参数
	if cType != "" && cType != ContentTypeFormUrlEncode &&
		cType != ContentTypeFormMultipart {
		return params
	}
	if ctx.Request == nil {
		return params
	}
	if ctx.Request.Method == "GET" {
		for k, v := range ctx.Request.URL.Query() {
			params[k] = v[0]
		}
		return params
	} else {
		err := ctx.Request.ParseForm()
		if err != nil {
			return params
		}
		for k, v := range ctx.Request.PostForm {
			params[k] = v[0]
		}
		for k, v := range ctx.Request.URL.Query() {
			params[k] = v[0]
		}
		return params
	}
}
