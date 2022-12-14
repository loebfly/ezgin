package call

import (
	"errors"
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/logs"
	"github.com/loebfly/ezgin/internal/nacos"
	"strings"
)

type formCall int

const Form = formCall(0)

func (receiver formCall) request(method, service, uri string, header, params map[string]string, files []grequests.FileUpload) (string, error) {
	return receiver.tryRequest(method, service, uri, header, params, files, true)
}

func (receiver formCall) tryRequest(method, service, uri string, header, params map[string]string, files []grequests.FileUpload, isFirstReq bool) (string, error) {
	if !isFirstReq {
		// 清除当前的服务缓存
		nacos.Enter.CleanServiceCache(service)
	}
	var url string
	var err error
	url, header, err = receiver.getReqUrlAndHeader(service, uri, header)
	if err != nil {
		logs.Enter.CError("CALL", "FORM - 获取{}服务地址失败:{}", service, err)
		return "", err
	}

	timeout := engine.MWTrace.GetCurTimeout()

	var resp *grequests.Response
	if files != nil {
		logs.Enter.CDebug("CALL",
			"FORM - FILE 微服务开始请求 -- url: {}, params: {}, header: {}",
			url, method, params, header)

		resp, err = grequests.Post(url, &grequests.RequestOptions{
			Data:               params,
			Files:              files,
			Headers:            header,
			RequestTimeout:     timeout,
			InsecureSkipVerify: true,
		})
		if err != nil {
			logs.Enter.CError("CALL",
				"FORM - FILE 微服务请求失败 -- url: {}, params: {}, header: {}, err: {}",
				url, params, header, err)
			if isFirstReq && strings.Contains(err.Error(), "connection refused") {
				return receiver.tryRequest(method, service, uri, header, params, files, false)
			}
			if strings.Contains(err.Error(), "dial tcp") {
				return "", errors.New("service unavailable")
			}
			return "", err
		}
	} else {
		logs.Enter.CDebug("CALL",
			"FORM - {} 微服务开始请求 -- url: {}, params: {}, header: {}",
			method, url, params, header)

		if method == "GET" {
			var options = &grequests.RequestOptions{
				Params:             params,
				Headers:            header,
				InsecureSkipVerify: true,
			}
			resp, err = grequests.Get(url, options)
		} else {
			var options = &grequests.RequestOptions{
				Data:               params,
				Headers:            header,
				InsecureSkipVerify: true,
			}
			resp, err = grequests.Post(url, options)
		}
		if err != nil {
			logs.Enter.CError("CALL",
				"FORM - {} 微服务请求失败 -- url: {}, params: {}, header: {}, err: {}",
				method, url, params, header, err)
			if isFirstReq && strings.Contains(err.Error(), "connection refused") {
				return receiver.tryRequest(method, service, uri, header, params, files, false)
			}
			if strings.Contains(err.Error(), "dial tcp") {
				return "", errors.New("service unavailable")
			}
			return "", err
		}
	}
	logs.Enter.CDebug("CALL",
		"FORM - {} 微服务请求响应 -- url: {}, params: {}, header: {}, resp: {}",
		method, url, params, header, resp.String())
	return resp.String(), nil

}

func (receiver formCall) getReqUrlAndHeader(service, uri string, header map[string]string) (string, map[string]string, error) {
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
	return url, header, nil
}
