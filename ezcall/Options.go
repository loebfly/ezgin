package ezcall

import (
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/engine"
	"time"
)

type OptionsProtocol interface{}

type FormOptions struct {
	Service string                 // 服务名
	Uri     string                 // 请求路径
	Method  engine.HttpMethod      // 请求方法
	timeout time.Duration          // 超时时间, 默认为30秒
	Header  map[string]string      // 请求头
	Files   []grequests.FileUpload // 文件
	Params  map[string]string      // 请求参数
}

func (receiver *FormOptions) SetTimeout(duration time.Duration) *FormOptions {
	receiver.timeout = duration
	return receiver
}

func (receiver *FormOptions) GetTimeout() time.Duration {
	return receiver.timeout
}

func (receiver *FormOptions) IsValid() bool {
	return receiver.Service != "" && receiver.Uri != "" && receiver.Method != ""
}

type JsonOptions struct {
	Service string            // 服务名
	Uri     string            // 请求路径
	Method  engine.HttpMethod // 请求方法
	timeout time.Duration     // 超时时间, 默认为30秒
	Header  map[string]string // 请求头
	Query   map[string]string // 请求参数
	JSON    any               // 请求体
}

func (receiver *JsonOptions) SetTimeout(duration time.Duration) *JsonOptions {
	receiver.timeout = duration
	return receiver
}

func (receiver *JsonOptions) GetTimeout() time.Duration {
	return receiver.timeout
}

func (receiver *JsonOptions) IsValid() bool {
	return receiver.Service != "" && receiver.Uri != "" && receiver.Method != ""
}

type RestfulOptions struct {
	Service string            // 服务名
	Uri     string            // 请求路径
	Method  engine.HttpMethod // 请求方法
	timeout time.Duration     // 超时时间, 默认为30秒
	Header  map[string]string // 请求头
	Path    map[string]string // 路径参数
	Query   map[string]string // 请求参数
	Body    any               // 请求体
}

func (receiver *RestfulOptions) SetTimeout(duration time.Duration) *RestfulOptions {
	receiver.timeout = duration
	return receiver
}

func (receiver *RestfulOptions) GetTimeout() time.Duration {
	return receiver.timeout
}

func (receiver *RestfulOptions) IsValid() bool {
	return receiver.Service != "" && receiver.Uri != "" && receiver.Method != ""
}
