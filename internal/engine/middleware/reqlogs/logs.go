package reqlogs

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/loebfly/ezgin/internal/engine/middleware/trace"
	"github.com/loebfly/ezgin/internal/logs"
	"io/ioutil"
	"strings"
	"time"
)

const (
	ContentTypeFormUrlEncode = "application/x-www-form-urlencoded"
	ContentTypeFormMultipart = "multipart/form-data"
)

func (receiver enter) Middleware(c *gin.Context) {
	// 不记录静态文件和根目录请求
	if strings.Contains(c.Request.RequestURI, "/docs") || c.Request.RequestURI == "/" {
		return
	}

	rWriter := &respWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}
	c.Writer = rWriter

	// 开始时间
	startTime := time.Now()
	reqTime := startTime.Format("2006-01-02 15:04:05.012")

	rawData, err := c.GetRawData()
	if err != nil {
		logs.Enter.CError("GIN", "GetRawData error:{}", err.Error())
	}

	// 关键点 重置请求体
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData))

	// 处理请求
	c.Next()

	var reqParams interface{}
	if strings.Contains(c.ContentType(), "application/json") {
		var params = make(map[string]interface{})
		err = json.Unmarshal(rawData, &reqParams)
		if err != nil {
			logs.Enter.CError("GIN", "reqParams json.unmarshal error:{}", err.Error())
		}
		reqParams = params
	} else if strings.Contains(c.ContentType(), "application/x-www-form-urlencoded") ||
		strings.Contains(c.ContentType(), "multipart/form-rawData") {
		reqParams = receiver.GetFormParams(c)
	} else {
		reqParams = string(rawData)
	}

	endTime := time.Now()
	respTime := endTime.Format("2006-01-02 15:04:05.012")

	var respParams = make(map[string]interface{})
	respStr := rWriter.body.String()
	if respStr != "" && respStr[0:1] == "{" {
		err = json.Unmarshal(rWriter.body.Bytes(), &respParams)
		if err != nil {
			logs.Enter.CError("GIN", "respParams json.Unmarshal error:{}", err.Error())
		}
	}

	ttl := int(endTime.UnixNano()/1e6 - startTime.UnixNano()/1e6)

	method := c.Request.Method
	contentType := c.ContentType()
	uri := c.Request.RequestURI

	logs.Enter.CError("GIN", "|{}|{}|{}|{}|{}ms", method, uri, c.ClientIP(), respTime, ttl)
	logs.Enter.CError("GIN", "请求参数:{}", reqParams)
	logs.Enter.CError("GIN", "接口返回:{}", respParams)

	ctx := ReqCtx{
		RequestId:   trace.Enter.GetCurReqId(),
		ReqTime:     reqTime,
		ReqParams:   reqParams,
		RespTime:    respTime,
		RespParams:  respParams,
		TTL:         ttl,
		Method:      method,
		ContentType: contentType,
		URI:         uri,
	}
	logChan <- ctx
}

func (receiver enter) GetFormParams(ctx *gin.Context) map[string]string {
	params := make(map[string]string)
	cType := ctx.ContentType()
	if cType != ContentTypeFormUrlEncode &&
		cType != ContentTypeFormMultipart {
		return params
	}
	if ctx.Request == nil {
		return params
	}
	if ctx.Request.Method == "GET" {
		for k, v := range ctx.Request.URL.Query() {
			params[k] = v[0]
		}
		return params
	} else {
		err := ctx.Request.ParseForm()
		if err != nil {
			return params
		}
		for k, v := range ctx.Request.PostForm {
			params[k] = v[0]
		}
		for k, v := range ctx.Request.URL.Query() {
			params[k] = v[0]
		}
		return params
	}
}
