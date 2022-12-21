package ezcall

import (
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/engine"
	"time"
)

type ReqWay int // 请求方式

const (
	ReqWayForm ReqWay = iota
	ReqWayJson
	ReqWayRestFul
)

type OptionsProtocol interface {
	GetMethod() engine.HttpMethod
	GetTimeout() time.Duration
	GetHeader() map[string]string
}

type baseOptions struct {
	Method  engine.HttpMethod
	Timeout time.Duration
	Header  map[string]string
}

func (receiver baseOptions) GetMethod() engine.HttpMethod {
	return receiver.Method
}

func (receiver baseOptions) GetTimeout() time.Duration {
	return receiver.Timeout
}

func (receiver baseOptions) GetHeader() map[string]string {
	return receiver.Header
}

type FormOptions struct {
	baseOptions
	Files  []grequests.FileUpload
	Params map[string]string
}

type JsonOptions struct {
	baseOptions
	Query map[string]string
	JSON  any
}

type RestFulOptions struct {
	baseOptions
	Path  map[string]string
	Query map[string]string
	Body  any
}
