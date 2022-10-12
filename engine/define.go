package engine

import "github.com/gin-gonic/gin"

type HttpMethod string

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

type HandlerFunc func(ctx *gin.Context) Result

type RecoveryFunc func(c *gin.Context, err interface{})

type Result struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    *Page       `json:"page"`
}

type Page struct {
	Count int `json:"count"`
	Index int `json:"index"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

const (
	ContentTypeFormUrlEncode = "application/x-www-form-urlencoded"
	ContentTypeFormMultipart = "multipart/form-data"
)

func GetFormParams(ctx *gin.Context) map[string]string {
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
