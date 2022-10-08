package call

import (
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/internal/engine/middleware/trace"
	"github.com/loebfly/ezgin/internal/logs"
	"github.com/loebfly/ezgin/internal/nacos"
)

type form int

const Form = form(0)

func (receiver form) post(service string, uri string, params map[string]string) (string, error) {
	return receiver.postWithHeader(service, uri, params, nil)
}

func (receiver form) postWithHeader(service string, uri string, params map[string]string, header map[string]string) (string, error) {
	return receiver.requestWithHeader("POST", service, uri, params, header)
}

func (receiver form) get(service, uri string, params map[string]string) (string, error) {
	return receiver.getWithHeader(service, uri, params, nil)
}

func (receiver form) getWithHeader(service, uri string, params map[string]string, header map[string]string) (string, error) {
	return receiver.requestWithHeader("GET", service, uri, params, header)
}

func (form) requestWithHeader(method, service, uri string, params map[string]string, header map[string]string) (string, error) {
	host, err := nacos.Enter.GetService(service)
	if err != nil {
		return "", err
	}
	url := host + uri
	logs.Enter.CDebug("CALL", "POST: {}, params: {}, headers: {}", url, params, header)
	if header == nil {
		header = make(map[string]string)
	}
	traceId := trace.Enter.GetCurReqId()
	if traceId != "" {
		header["X-Request-Id"] = traceId
	}
	//header["X-Lang"] = xlang.GetCurrentLanguage()
	var resp *grequests.Response
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
	return resp.String(), nil

}

func (receiver form) file(service string, uri string, params map[string]string, files []grequests.FileUpload) (string, error) {
	return receiver.fileWithHeader(service, uri, params, files, nil)
}

func (form) fileWithHeader(service string, uri string, params map[string]string, files []grequests.FileUpload, header map[string]string) (string, error) {
	return "", nil
}
