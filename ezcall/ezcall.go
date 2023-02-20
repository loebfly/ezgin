package ezcall

import (
	"encoding/json"
	"errors"
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/ezlogs"
	"github.com/loebfly/ezgin/internal/call"
	"reflect"
)

/********* 通用 *********/

// Request 发起请求, service 服务名, uri 请求地址, option 请求参数(FormOptions, JsonOptions, RestfulOptions)
func Request(option OptionsProtocol) (*grequests.Response, error) {
	if option == nil {
		return nil, errors.New("option is nil")
	}
	// 判断option 是否为指针地址
	if reflect.ValueOf(option).Kind() != reflect.Ptr {
		return nil, errors.New("options 需要传入指针地址")
	}
	if reflect.ValueOf(option).IsNil() {
		return nil, errors.New("option is nil")
	}
	// 判断option类型
	typeStr := reflect.TypeOf(option).String()
	if typeStr == "*ezcall.FormOptions" {
		formOptions := option.(*FormOptions)
		if !formOptions.IsValid() {
			return nil, errors.New("options参数不完整")
		}
		return call.Form.Request(string(formOptions.Method), formOptions.Service, formOptions.Uri, formOptions.Header, formOptions.Params, formOptions.Files)
	} else if typeStr == "*ezcall.JsonOptions" {
		jsonOption := option.(*JsonOptions)
		if !jsonOption.IsValid() {
			return nil, errors.New("options参数不完整")
		}
		return call.Json.Request(jsonOption.Method, jsonOption.Service, jsonOption.Uri, jsonOption.Header, jsonOption.Query, jsonOption.JSON)
	} else if typeStr == "*ezcall.RestfulOptions" {
		restfulOption := option.(*RestfulOptions)
		if !restfulOption.IsValid() {
			return nil, errors.New("options参数不完整")
		}
		return call.Restful.Request(restfulOption.Method, restfulOption.Service, restfulOption.Uri, restfulOption.Header, restfulOption.Path, restfulOption.Query, restfulOption.Body)
	} else {
		return nil, errors.New("option is not *FormOptions/*JsonOptions/*RestfulOptions")
	}
}

func RequestToResult[D any](option OptionsProtocol) engine.Result[D] {
	return toResult[D](Request(option))
}

/******* Form ******* */

func FormPostToResult[D any](service, uri string, params map[string]string) engine.Result[D] {
	return FormPostWithHeaderToResult[D](service, uri, nil, params)
}

func FormPostWithHeaderToResult[D any](service, uri string, header, params map[string]string) engine.Result[D] {
	return toResult[D](FormPostWithHeader(service, uri, header, params))
}

func FormPost(service, uri string, params map[string]string) (resp *grequests.Response, err error) {
	return FormPostWithHeader(service, uri, nil, params)
}

func FormPostWithHeader(service, uri string, header, params map[string]string) (resp *grequests.Response, err error) {
	return call.Form.Request("POST", service, uri, header, params, nil)
}

func FormGetToResult[D any](service, uri string, params map[string]string) engine.Result[D] {
	return FormGetWithHeaderToResult[D](service, uri, nil, params)
}

func FormGetWithHeaderToResult[D any](service, uri string, header, params map[string]string) engine.Result[D] {
	return toResult[D](FormGetWithHeader(service, uri, header, params))
}

func FormGet(service, uri string, params map[string]string) (resp *grequests.Response, err error) {
	return FormGetWithHeader(service, uri, params, nil)
}

func FormGetWithHeader(service, uri string, header, params map[string]string) (resp *grequests.Response, err error) {
	return call.Form.Request("GET", service, uri, header, params, nil)
}

func FormFileToResult[D any](service string, uri string, params map[string]string, files []grequests.FileUpload) engine.Result[D] {
	return FormFileWithHeaderToResult[D](service, uri, nil, params, files)
}

func FormFileWithHeaderToResult[D any](service string, uri string, header, params map[string]string, files []grequests.FileUpload) engine.Result[D] {
	return toResult[D](FormFileWithHeader(service, uri, header, params, files))
}

func FormFile(service string, uri string, params map[string]string, files []grequests.FileUpload) (resp *grequests.Response, err error) {
	return FormFileWithHeader(service, uri, params, nil, files)
}

func FormFileWithHeader(service string, uri string, header, params map[string]string, files []grequests.FileUpload) (resp *grequests.Response, err error) {
	return call.Form.Request("POST", service, uri, header, params, files)
}

/******* Json ******* */

func JsonToResult[D any](method engine.HttpMethod, service string, uri string, query map[string]string, body any) engine.Result[D] {
	return JsonWithHeaderToResult[D](method, service, uri, nil, query, body)
}

func JsonWithHeaderToResult[D any](method engine.HttpMethod, service string, uri string, header, query map[string]string, body any) engine.Result[D] {
	return toResult[D](JsonWithHeader(method, service, uri, header, query, body))
}

func Json(method engine.HttpMethod, service string, uri string, query map[string]string, body any) (resp *grequests.Response, err error) {
	return JsonWithHeader(method, service, uri, nil, query, body)
}

func JsonWithHeader(method engine.HttpMethod, service string, uri string, header, query map[string]string, body any) (resp *grequests.Response, err error) {
	return call.Json.Request(method, service, uri, header, query, body)
}

/******* Restful ******* */

func RestfulToResult[D any](method engine.HttpMethod, service string, uri string, path, query map[string]string, body any) engine.Result[D] {
	return RestfulWithHeaderToResult[D](method, service, uri, path, nil, query, body)
}

func RestfulWithHeaderToResult[D any](method engine.HttpMethod, service string, uri string, path, header, query map[string]string, body any) engine.Result[D] {
	return toResult[D](RestfulWithHeader(method, service, uri, path, header, query, body))
}

func Restful(method engine.HttpMethod, service string, uri string, path, query map[string]string, body any) (resp *grequests.Response, err error) {
	return RestfulWithHeader(method, service, uri, path, nil, query, body)
}

func RestfulWithHeader(method engine.HttpMethod, service string, uri string, path, header, query map[string]string, body any) (resp *grequests.Response, err error) {
	return call.Restful.Request(method, service, uri, path, header, query, body)
}

// toResult 将字符串转换为Result
func toResult[D any](resp *grequests.Response, err error) engine.Result[D] {
	if err != nil {
		return engine.Result[D]{
			Status:  engine.ErrorCodeServiceUnavailable,
			Message: err.Error(),
		}
	}
	if resp == nil || resp.String() == "" {
		ezlogs.Error("resp:{}, 返回结果为空", resp)
		return engine.Result[D]{
			Status:  engine.ErrorCodeResUnmarshalFailure,
			Message: "response is nil",
		}
	}
	var result engine.Result[D]
	err = json.Unmarshal([]byte(resp.String()), &result)
	if err != nil {
		ezlogs.Error("resp:{}, 返回结果序列化失败: {}", resp, err.Error())
		return engine.Result[D]{
			Status:  engine.ErrorCodeResUnmarshalFailure,
			Message: "response unmarshal failure",
		}
	} else {
		return result
	}
}
