package call

import (
	"encoding/json"
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/engine"
)

type enter int

const Enter = enter(0)

/******* Form ******* */

func (receiver enter) FormPostToResult(service, uri string, params map[string]string) engine.Result {
	return receiver.FormPostWithHeaderToResult(service, uri, nil, params)
}

func (receiver enter) FormPostWithHeaderToResult(service, uri string, header, params map[string]string) engine.Result {
	return receiver.toResult(receiver.FormPostWithHeader(service, uri, header, params))
}

func (receiver enter) FormPost(service, uri string, params map[string]string) (string, error) {
	return receiver.FormPostWithHeader(service, uri, nil, params)
}

func (receiver enter) FormPostWithHeader(service, uri string, header, params map[string]string) (string, error) {
	return Form.request("POST", service, uri, header, params, nil)
}

func (receiver enter) FormGetToResult(service, uri string, params map[string]string) engine.Result {
	return receiver.FormGetWithHeaderToResult(service, uri, nil, params)
}

func (receiver enter) FormGetWithHeaderToResult(service, uri string, header, params map[string]string) engine.Result {
	return receiver.toResult(receiver.FormGetWithHeader(service, uri, header, params))
}

func (receiver enter) FormGet(service, uri string, params map[string]string) (string, error) {
	return receiver.FormGetWithHeader(service, uri, params, nil)
}

func (receiver enter) FormGetWithHeader(service, uri string, header, params map[string]string) (string, error) {
	return Form.request("GET", service, uri, header, params, nil)
}

func (receiver enter) FormFileToResult(service string, uri string, params map[string]string, files []grequests.FileUpload) engine.Result {
	return receiver.FormFileWithHeaderToResult(service, uri, nil, params, files)
}

func (receiver enter) FormFileWithHeaderToResult(service string, uri string, header, params map[string]string, files []grequests.FileUpload) engine.Result {
	return receiver.toResult(receiver.FormFileWithHeader(service, uri, header, params, files))
}

func (receiver enter) FormFile(service string, uri string, params map[string]string, files []grequests.FileUpload) (string, error) {
	return receiver.FormFileWithHeader(service, uri, params, nil, files)
}

func (receiver enter) FormFileWithHeader(service string, uri string, header, params map[string]string, files []grequests.FileUpload) (string, error) {
	return Form.request("POST", service, uri, header, params, files)
}

/******* Json ******* */

func (receiver enter) JsonToResult(method engine.HttpMethod, service string, uri string, query map[string]string, body interface{}) engine.Result {
	return receiver.JsonWithHeaderToResult(method, service, uri, nil, query, body)
}

func (receiver enter) JsonWithHeaderToResult(method engine.HttpMethod, service string, uri string, header, query map[string]string, body interface{}) engine.Result {
	return receiver.toResult(receiver.JsonWithHeader(method, service, uri, header, query, body))
}

func (receiver enter) Json(method engine.HttpMethod, service string, uri string, query map[string]string, body interface{}) (string, error) {
	return receiver.JsonWithHeader(method, service, uri, nil, query, body)
}

func (receiver enter) JsonWithHeader(method engine.HttpMethod, service string, uri string, header, query map[string]string, body interface{}) (string, error) {
	return Json.request(method, service, uri, header, query, body)
}

/******* Restful ******* */

func (receiver enter) RestfulToResult(method engine.HttpMethod, service string, uri string, path, query map[string]string, body interface{}) engine.Result {
	return receiver.RestfulWithHeaderToResult(method, service, uri, path, nil, query, body)
}

func (receiver enter) RestfulWithHeaderToResult(method engine.HttpMethod, service string, uri string, path, header, query map[string]string, body interface{}) engine.Result {
	return receiver.toResult(receiver.RestfulWithHeader(method, service, uri, path, header, query, body))
}

func (receiver enter) Restful(method engine.HttpMethod, service string, uri string, path, query map[string]string, body interface{}) (string, error) {
	return receiver.RestfulWithHeader(method, service, uri, path, nil, query, body)
}

func (receiver enter) RestfulWithHeader(method engine.HttpMethod, service string, uri string, path, header, query map[string]string, body interface{}) (string, error) {
	return Restful.request(method, service, uri, path, header, query, body)
}

// toResult 将字符串转换为Result
func (receiver enter) toResult(resp string, err error) engine.Result {
	if err != nil {
		return engine.Result{
			Status:  engine.ErrorCodeServiceUnavailable,
			Message: err.Error(),
		}
	}
	var result engine.Result
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return engine.Result{
			Status:  engine.ErrorCodeResUnmarshalFailure,
			Message: err.Error(),
		}
	} else {
		return result
	}
}
