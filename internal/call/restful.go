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

func (receiver restfulCall) request(method define.HttpMethod, service, uri string, path, header, query map[string]string, body interface{}) (string, error) {
	var url string
	var err error
	url, header, err = receiver.getReqUrlAndHeader(service, uri, path, header)
	if err != nil {
		logs.Enter.CError("CALL", "RESTFUL - 获取{}服务地址失败:{}", service, err)
		return "", err
	}
	logs.Enter.CDebug("CALL", "RESTFUL - {}微服务请求开始 -- url: {}, headers: {}, query: {}, body: {}", method, url, header, query, body)
	var resp *grequests.Response
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
		return "", errors.New("RESTFUL - 不支持的请求方法")
	}
	if err != nil {
		logs.Enter.CError("CALL", "RESTFUL - 请求{}微服务失败:{}", service, err)
		return "", err
	}
	logs.Enter.CDebug("CALL", "RESTFUL - {}微服务请求结束 -- url: {}, headers: {}, query: {}, body:{}, resp: {}", method, url, header, query, body, resp.String())
	return resp.String(), nil
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
	if header == nil {
		header = make(map[string]string)
	}
	reqId := engine.MWTrace.GetCurReqId()
	if reqId != "" {
		header["X-Request-Id"] = reqId
	}
	realIp := engine.MWTrace.GetCurClientIP()
	if realIp != "" {
		header["X-Real-Ip"] = realIp
	}
	userAgent := engine.MWTrace.GetCurUserAgent()
	if userAgent != "" {
		header["X-User-Agent"] = userAgent
	}
	xLang := engine.MWXLang.GetCurXLang()
	if xLang != "" {
		header["X-Lang"] = xLang
	}
	header["Content-Type"] = "application/json"
	return url, header, nil
}
