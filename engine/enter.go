package engine

import "github.com/gin-gonic/gin"

type enter int

const Enter = enter(0)

type HttpMethod string

const (
	Any     HttpMethod = "ANY"
	Get     HttpMethod = "GET"
	Head    HttpMethod = "HEAD"
	Post    HttpMethod = "POST"
	Put     HttpMethod = "PUT"
	Patch   HttpMethod = "PATCH" // RFC 5789
	Delete  HttpMethod = "DELETE"
	Connect HttpMethod = "CONNECT"
	Options HttpMethod = "OPTIONS"
	Trace   HttpMethod = "TRACE"
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
