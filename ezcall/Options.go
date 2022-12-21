package ezcall

import (
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/engine"
	"time"
)

type OptionsProtocol interface {
	GetMethod() engine.HttpMethod
	GetTimeout() time.Duration
	GetHeader() map[string]string
}

type FormOptions struct {
	Method  engine.HttpMethod
	Timeout time.Duration
	Header  map[string]string
	Files   []grequests.FileUpload
	Params  map[string]string
}

func (receiver FormOptions) GetMethod() engine.HttpMethod {
	return receiver.Method
}

func (receiver FormOptions) GetTimeout() time.Duration {
	return receiver.Timeout
}

func (receiver FormOptions) GetHeader() map[string]string {
	return receiver.Header
}

type JsonOptions struct {
	Method  engine.HttpMethod
	Timeout time.Duration
	Header  map[string]string
	Query   map[string]string
	JSON    any
}

func (receiver JsonOptions) GetMethod() engine.HttpMethod {
	return receiver.Method
}

func (receiver JsonOptions) GetTimeout() time.Duration {
	return receiver.Timeout
}

func (receiver JsonOptions) GetHeader() map[string]string {
	return receiver.Header
}

type RestfulOptions struct {
	Method  engine.HttpMethod
	Timeout time.Duration
	Header  map[string]string
	Path    map[string]string
	Query   map[string]string
	Body    any
}

func (receiver RestfulOptions) GetMethod() engine.HttpMethod {
	return receiver.Method
}

func (receiver RestfulOptions) GetTimeout() time.Duration {
	return receiver.Timeout
}

func (receiver RestfulOptions) GetHeader() map[string]string {
	return receiver.Header
}
