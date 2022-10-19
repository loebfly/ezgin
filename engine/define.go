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

type MiddlewareFunc func(c *gin.Context)

type HandlerFunc func(c *gin.Context) Result

type RecoveryFunc func(c *gin.Context, err interface{})

type Result struct {
	Status  int         `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	Page    *Page       `json:"page"`
}

type Page struct {
	Count int `json:"count"`
	Index int `json:"index"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

func ErrorRes(status int, message string) Result {
	return Result{
		Status:  status,
		Message: message,
	}
}

func SuccessRes(data interface{}, message ...string) Result {
	msg := "success"
	if len(message) > 0 {
		msg = message[0]
	}
	return Result{
		Status:  1,
		Message: msg,
		Data:    data,
	}
}

func SuccessPageRes(data interface{}, page Page, message ...string) Result {
	msg := "success"
	if len(message) > 0 {
		msg = message[0]
	}
	return Result{
		Status:  1,
		Message: msg,
		Data:    data,
		Page:    &page,
	}
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
