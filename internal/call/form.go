package call

import (
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/internal/engine"
	"github.com/loebfly/ezgin/internal/logs"
	"github.com/loebfly/ezgin/internal/nacos"
)

type form int

const Form = form(0)

func (receiver form) request(method, service, uri string, header, params map[string]string, files []grequests.FileUpload) (string, error) {
	var url string
	var err error
	url, header, err = receiver.getReqUrlAndHeader(service, uri, header)
	if err != nil {
		return "", err
	}
	logs.Enter.CDebug("CALL",
		"Nacos -- {}微服务调用 -- url: {}, params: {}, headers: {}",
		method, url, params, header)
	var resp *grequests.Response
	if files != nil {
		resp, err = grequests.Post(url, &grequests.RequestOptions{
			Data:               params,
			Files:              files,
			Headers:            header,
			InsecureSkipVerify: true,
		})
		if err != nil {
			return "", err
		}
	} else {
		if method == "GET" {
			resp, err = grequests.Get(url, &grequests.RequestOptions{
				Params:             params,
				Headers:            header,
				InsecureSkipVerify: true,
			})
			if err != nil {
				return "", err
			}
		} else {
			resp, err = grequests.Post(url, &grequests.RequestOptions{
				Params:             params,
				Headers:            header,
				InsecureSkipVerify: true,
			})
		}
	}
	logs.Enter.CDebug("CALL",
		"Nacos -- {} 微服务响应 -- url: {} method: {}, params: {}, headers: {}, resp: {}",
		method, url, params, header, resp.String())
	return resp.String(), nil

}

func (receiver form) getReqUrlAndHeader(service, uri string, header map[string]string) (string, map[string]string, error) {
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
