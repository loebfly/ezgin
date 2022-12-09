package call

import (
	"errors"
	"github.com/levigross/grequests"
	define "github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/ezlogs"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/nacos"
)

type jsonCall int

const Json = jsonCall(0)

func (receiver jsonCall) Request(method define.HttpMethod, service, uri string, header, query map[string]string, body any) (resp *grequests.Response, err error) {
	var url string
	url, header, err = receiver.getReqUrlAndHeader(service, uri, header)
	if err != nil {
		ezlogs.CError("CALL", "JSON - 获取{}服务地址失败:{}", service, err)
		return
	}
	ezlogs.CDebug("CALL",
		"JSON - {}微服务请求开始 -- url: {}, headers: {}, query: {}, body: {}",
		method, url, header, query, body)

	var options = &grequests.RequestOptions{
		Params:             query,
		Headers:            header,
		InsecureSkipVerify: true,
		JSON:               body,
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
		err = errors.New("不支持的请求方法")
		return

	}
	if err != nil {
		ezlogs.CError("CALL",
			"JSON - {}微服务请求失败 -- url: {}, params: {}, headers: {}, err: {}",
			method, url, query, header, err)
		return
	}
	ezlogs.CDebug("CALL",
		"JSON - {} 微服务请求响应 -- url: {} method: {}, params: {}, headers: {}, resp: {}",
		method, url, query, header, resp.String())
	return
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
	header["Content-Type"] = "application/json"
	return url, header, nil
}
