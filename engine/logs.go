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
	ReqParams   any               `json:"requestParam" bson:"requestParam"`   // 请求参数
	RespParams  any               `json:"responseMap" bson:"responseMap"`     // 响应参数
}

func (ctx ReqCtx) GetRealMgoTag(tag string) string {
	// 可配置变量: {header:xxx}，表示从请求头中获取key为xxx对应的值替换变量
	// 假设请求头中有一个key为: X-Request-Id，值为: abc123
	// 配置文件中的变量为: {header:X-Request-Id}
	// 则最终的变量值为: abc123
	if !strings.Contains(tag, "{header:") {
		return tag
	}

	// 获取请求头中的key
	key := strings.Split(strings.Split(tag, "{header:")[1], "}")[0]
	if key == "" {
		return tag
	}

	// 获取请求头中的值
	value := ctx.ReqHeaders[key]
	if value == "" {
		return tag
	}

	// 替换变量
	realTag := strings.Replace(tag, "{header:"+key+"}", value, -1)
	return ctx.GetRealMgoTag(realTag)
}

func (ctx ReqCtx) GetRealMgoTable(table string) string {
	// 日志表名, 默认为${App.Name}APIRequestLogs,
	// 可配置变量: {header:xxx}，表示从请求头中获取key为xxx对应的值替换变量
	// 假设请求头中有一个key为: X-Request-Id，值为: abc123
	// 效果为:{header:X-Request-Id}_APIRequestLogs -> abc123_APIRequestLogs
	if !strings.Contains(table, "{header:") {
		return table
	}

	// 获取请求头中的key
	key := strings.Split(strings.Split(table, "{header:")[1], "}")[0]
	if key == "" {
		return table
	}

	// 获取请求头中的值
	value := ctx.ReqHeaders[key]
	if value == "" {
		return table
	}

	// 替换变量
	realTable := strings.Replace(table, "{header:"+key+"}", value, -1)
	return ctx.GetRealMgoTable(realTable)
}

func (ctx ReqCtx) GetRealKafkaTopic(topic string) string {
	// kafka 消息主题, 默认为${App.Name}, 多个主题用逗号分隔, - 表示不开启
	// 可配置变量: {header:xxx}，表示从请求头中获取key为xxx对应的值替换变量
	// 假设请求头中有一个key为: X-Request-Id，值为: abc123
	// 效果为:{header:X-Request-Id}_Topic -> abc123_Topic

	if !strings.Contains(topic, "{header:") {
		return topic
	}

	// 获取请求头中的key
	key := strings.Split(strings.Split(topic, "{header:")[1], "}")[0]
	if key == "" {
		return topic
	}

	// 获取请求头中的值
	value := ctx.ReqHeaders[key]
	if value == "" {
		return topic
	}

	// 替换变量
	realTopic := strings.Replace(topic, "{header:"+key+"}", value, -1)
	return ctx.GetRealKafkaTopic(realTopic)
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
