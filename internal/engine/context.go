package engine

import "github.com/gin-gonic/gin"

func (receiver control) GetFormParams(ctx *gin.Context) map[string]string {
	params := make(map[string]string)
	cType := ctx.ContentType()
	if cType != ContentTypeFormUrlEncode &&
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
