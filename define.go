package ezgin

import "github.com/gin-gonic/gin"

type HttpMethod string

const (
	HttpMethodAny     HttpMethod = "ANY"
	HttpMethodGet     HttpMethod = "GET"
	HttpMethodHead    HttpMethod = "HEAD"
	HttpMethodPost    HttpMethod = "POST"
	HttpMethodPut     HttpMethod = "PUT"
	HttpMethodPatch   HttpMethod = "PATCH"
	HttpMethodDelete  HttpMethod = "DELETE"
	HttpMethodConnect HttpMethod = "CONNECT"
	HttpMethodOptions HttpMethod = "OPTIONS"
	HttpMethodTrace   HttpMethod = "TRACE"
)

type GinHandlerFunc func(ctx *gin.Context) GinResult

type GinResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    struct {
		Count int `json:"count"`
		Index int `json:"index"`
		Size  int `json:"size"`
		Total int `json:"total"`
	} `json:"page"`
}
