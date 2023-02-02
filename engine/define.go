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

type HandlerFunc func(c *gin.Context) Result[any]

type RecoveryFunc func(c *gin.Context, err any)

type Result[D any] struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
	Data    D      `json:"data"`
	Page    *Page  `json:"page"`
}

// ToAnyRes 转换为Result[any]
func (receiver Result[D]) ToAnyRes() Result[any] {
	if receiver.Status != 1 {
		return Result[any]{
			Status:  receiver.Status,
			Message: receiver.Message,
		}
	}
	return Result[any]{
		Status:  receiver.Status,
		Message: receiver.Message,
		Data:    receiver.Data,
		Page:    receiver.Page,
	}
}

// FreeConvResDataType 自由将Result[From]转换为Result[To]
func FreeConvResDataType[From, To any](from Result[From]) Result[To] {
	res := from.ToAnyRes()
	return Result[To]{
		Status:  res.Status,
		Message: res.Message,
		Data:    res.Data.(To),
		Page:    res.Page,
	}
}

// ConvAnyResDataType 将Result[any]转换为Result[D]
func ConvAnyResDataType[To any](from Result[any]) Result[To] {
	return Result[To]{
		Status:  from.Status,
		Message: from.Message,
		Data:    from.Data.(To),
		Page:    from.Page,
	}
}

type Page struct {
	Count int `json:"count"`
	Index int `json:"index"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

func ErrorRes(status int, message string) Result[any] {
	return Result[any]{
		Status:  status,
		Message: message,
	}
}

func SuccessRes[D any](data D, message ...string) Result[D] {
	msg := "success"
	if len(message) > 0 {
		msg = message[0]
	}
	return Result[D]{
		Status:  1,
		Message: msg,
		Data:    data,
	}
}

func SuccessPageRes[D any](data D, page Page, message ...string) Result[D] {
	msg := "success"
	if len(message) > 0 {
		msg = message[0]
	}
	return Result[D]{
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
