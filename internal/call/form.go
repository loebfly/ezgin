package call

import (
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/logs"
	"github.com/loebfly/ezgin/internal/nacos"
)

type formCall int

const Form = formCall(0)

func (receiver formCall) request(method, service, uri string, header, params map[string]string, files []grequests.FileUpload) (string, error) {
	var url string
	var err error
	url, header, err = receiver.getReqUrlAndHeader(service, uri, header)
	if err != nil {
		logs.Enter.CError("CALL", "FORM - 获取{}服务地址失败:{}", service, err)
		return "", err
	}

	var resp *grequests.Response
	if files != nil {
		logs.Enter.CDebug("CALL",
			"FORM - FILE微服务开始请求 -- url: {}, params: {}, headers: {}",
			method, url, params, header)
		resp, err = grequests.Post(url, &grequests.RequestOptions{
			Data:               params,
			Files:              files,
			Headers:            header,
			InsecureSkipVerify: true,
		})
		if err != nil {
			logs.Enter.CError("CALL", "FORM - FILE微服务请求失败 -- url: {}, params: {}, headers: {}, err: {}", method, url, params, header, err)
			return "", err
		}
	} else {
		logs.Enter.CDebug("CALL",
			"FORM - {}微服务开始请求 -- url: {}, params: {}, headers: {}",
			method, url, params, header)
		var options = &grequests.RequestOptions{
			Data:               params,
			Headers:            header,
			InsecureSkipVerify: true,
		}
		if method == "GET" {
			resp, err = grequests.Get(url, options)
		} else {
			resp, err = grequests.Post(url, options)
		}
		if err != nil {
			logs.Enter.CError("CALL", "FORM - {}微服务请求失败 -- url: {}, params: {}, headers: {}, err: {}", method, url, params, header, err)
			return "", err
		}
	}
	logs.Enter.CDebug("CALL",
		"FORM - {} 微服务请求响应 -- url: {} method: {}, params: {}, headers: {}, resp: {}",
		method, url, params, header, resp.String())
	return resp.String(), nil

}

func (receiver formCall) getReqUrlAndHeader(service, uri string, header map[string]string) (string, map[string]string, error) {
	host, err := nacos.Enter.GetService(service)
	if err != nil {
		return "", nil, err
	}
	url := host + uri
	if header == nil {
		header = make(map[string]string)
	}
	reqId := engine.Enter.GetCurReqId()
	if reqId != "" {
		header["X-Request-Id"] = reqId
	}
	xLang := engine.Enter.GetCurXLang()
	if xLang != "" {
		header["X-Lang"] = xLang
	}
	return url, header, nil
}
