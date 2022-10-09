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
