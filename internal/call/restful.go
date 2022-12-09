package call

import (
	"errors"
	"github.com/levigross/grequests"
	define "github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/logs"
	"github.com/loebfly/ezgin/internal/nacos"
	"strings"
)

type restfulCall int

const Restful = restfulCall(0)

func (receiver restfulCall) Request(method define.HttpMethod, service, uri string, path, header, query map[string]string, body any) (resp *grequests.Response, err error) {
	var url string
	url, header, err = receiver.getReqUrlAndHeader(service, uri, path, header)
	if err != nil {
		logs.Enter.CError("CALL", "RESTFUL - 获取{}服务地址失败:{}", service, err)
		return
	}
	logs.Enter.CDebug("CALL", "RESTFUL - {}微服务请求开始 -- url: {}, headers: {}, query: {}, body: {}", method, url, header, query, body)
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
		logs.Enter.CError("CALL", "RESTFUL - 不支持的请求方法:{}", method)
		err = errors.New("不支持的请求方法")
		return
	}
	if err != nil {
		logs.Enter.CError("CALL", "RESTFUL - 请求{}微服务失败:{}", service, err)
		return
	}
	logs.Enter.CDebug("CALL", "RESTFUL - {}微服务请求结束 -- url: {}, headers: {}, query: {}, body:{}, resp: {}", method, url, header, query, body, resp.String())
	return resp, nil
}

func (receiver restfulCall) getReqUrlAndHeader(service, uri string, path, header map[string]string) (string, map[string]string, error) {
	host, err := nacos.Enter.GetService(service)
	if err != nil {
		return "", nil, err
	}

	for k, v := range path {
		uri = strings.ReplaceAll(uri, "{"+k+"}", v)
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
