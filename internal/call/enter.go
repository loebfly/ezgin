package call

import "github.com/levigross/grequests"

type enter int

const Enter = enter(0)

/******* Form ******* */

func (receiver enter) FormPost(service string, uri string, params map[string]string) (string, error) {
	return receiver.FormPostWithHeader(service, uri, nil, params)
}

func (receiver enter) FormPostWithHeader(service string, uri string, header, params map[string]string) (string, error) {
	return Form.request("POST", service, uri, header, params, nil)
}

func (receiver enter) FormGet(service, uri string, params map[string]string) (string, error) {
	return receiver.FormGetWithHeader(service, uri, params, nil)
}

func (receiver enter) FormGetWithHeader(service, uri string, header, params map[string]string) (string, error) {
	return Form.request("GET", service, uri, header, params, nil)
}

func (receiver enter) FormFile(service string, uri string, params map[string]string, files []grequests.FileUpload) (string, error) {
	return receiver.FormFileWithHeader(service, uri, params, nil, files)
}

func (receiver enter) FormFileWithHeader(service string, uri string, header, params map[string]string, files []grequests.FileUpload) (string, error) {
	return Form.request("POST", service, uri, header, params, files)
}

/******* Json ******* */
