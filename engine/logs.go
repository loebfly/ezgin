package engine

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type ReqCtx struct {
	Id          bson.ObjectId     `json:"-" bson:"_id"`
	ReqTime     string            `json:"time" bson:"time"`                   // 请求时间
	RequestId   string            `json:"requestId" bson:"requestId"`         // 请求id
	RespTime    string            `json:"responseTime" bson:"responseTime"`   // 响应时间
	TTL         int               `json:"ttl" bson:"ttl"`                     // 请求耗时
	AppName     string            `json:"appName" bson:"appName"`             // 应用名称
	Method      string            `json:"method" bson:"method"`               // 请求方法
	ContentType string            `json:"contentType" bson:"contentType"`     // 请求类型
	URI         string            `json:"uri" bson:"uri"`                     // 请求URI
	ClientIP    string            `json:"clientIP" bson:"clientIP"`           // 客户端IP
	ReqHeaders  map[string]string `json:"requestHeader" bson:"requestHeader"` // 请求头
	ReqParams   interface{}       `json:"requestParam" bson:"requestParam"`   // 请求参数
	RespParams  interface{}       `json:"responseMap" bson:"responseMap"`     // 响应参数
}

func (ctx ReqCtx) ToJson() string {
	b, err := json.Marshal(ctx)
	if err != nil {
		return ""
	}
	result := string(b)
	result = strings.Replace(result, "\\u003c", "<", -1)
	result = strings.Replace(result, "\\u003e", ">", -1)
	result = strings.Replace(result, "\\u0026", "&", -1)
	return result
}
