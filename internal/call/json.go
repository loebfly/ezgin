package call

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/levigross/grequests"
	define "github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/ezlogs"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/nacos"
	"strings"
)

type jsonCall int

const Json = jsonCall(0)

func (receiver jsonCall) Request(method define.HttpMethod, service, uri string, header, query map[string]string, JSON any) (resp *grequests.Response, err error) {
	return receiver.tryRequest(method, service, uri, header, query, JSON, true)
}

func (receiver jsonCall) tryRequest(method define.HttpMethod, service, uri string, header, query map[string]string, JSON any, isFirstReq bool) (resp *grequests.Response, err error) {
	if !isFirstReq {
		// 清除当前的服务缓存
		nacos.Enter.CleanServiceCache(service)
	}
	var url string
	url, header, err = receiver.getReqUrlAndHeader(service, uri, header)
	if err != nil {
		ezlogs.CError("CALL", "JSON - 获取{}服务地址失败:{}", service, err)
		return nil, err
	}
	ezlogs.CDebug("CALL",
		"JSON - {}微服务请求开始 -- url: {}, headers: {}, query: {}, JSON: {}",
		method, url, header, query, JSON)

	timeout := engine.MWTrace.GetCurTimeout()
	var options = &grequests.RequestOptions{
		Params:             query,
		Headers:            header,
		InsecureSkipVerify: true,
		RequestTimeout:     timeout,
		JSON:               JSON,
	}
	switch method {
	case define.Get:
		resp, err = grequests.Get(url, options)
	case define.Post:
		resp, err = grequests.Post(url, options)
	case define.Put:
		resp, err = grequests.Put(url, options)
	case define.Delete:
		resp, err = grequests.Delete(url, options)
	case define.Options:
		resp, err = grequests.Options(url, options)
	case define.Head:
		resp, err = grequests.Head(url, options)
	case define.Patch:
		resp, err = grequests.Patch(url, options)
	default:
		return nil, errors.New("不支持的请求方法")

	}
	if err != nil {
		ezlogs.CDebug("CALL",
			"JSON - {}微服务请求失败 -- url: {}, headers: {}, query: {}, JSON: {}, err: {}",
			method, url, header, query, JSON, err.Error())
		if isFirstReq && strings.Contains(err.Error(), "connection refused") {
			return receiver.tryRequest(method, service, uri, header, query, JSON, false)
		}
		if strings.Contains(err.Error(), "dial tcp") {
			return nil, errors.New("service unavailable")
		}
		return nil, err
	}
	ezlogs.CDebug("CALL",
		"JSON - {}微服务请求响应 -- url: {}, headers: {}, query: {}, JSON: {}, resp: {}",
		method, url, header, query, JSON, resp.String())
	return resp, nil
}

func (receiver jsonCall) getReqUrlAndHeader(service, uri string, header map[string]string) (string, map[string]string, error) {
	host, err := nacos.Enter.GetService(service)
	if err != nil {
		return "", nil, err
	}
	url := host + uri
	traceHeader := engine.MWTrace.GetCurHeader()
	if header == nil {
		header = traceHeader
	} else {
		for k, v := range traceHeader {
			if _, ok := header[k]; !ok {
				header[k] = v
			}
		}
	}
	header["Content-Type"] = gin.MIMEJSON
	return url, header, nil
}
