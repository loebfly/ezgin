package reqlogs

import "gopkg.in/mgo.v2/bson"

type ReqCtx struct {
	Id          bson.ObjectId       `bson:"_id"`
	RequestId   string              `bson:"request_id"`      // 请求id
	ReqTime     string              `bson:"request_time"`    // 请求时间
	ReqHeaders  map[string][]string `bson:"request_headers"` // 请求头
	ReqParams   interface{}         `bson:"request_params"`  // 请求参数
	RespTime    string              `bson:"response_time"`   // 响应时间
	RespParams  interface{}         `bson:"response_params"` // 响应参数
	TTL         int                 `bson:"ttl"`             // 请求耗时
	Method      string              `bson:"method"`          // 请求方法
	ContentType string              `bson:"content_type"`    // 请求类型
	URI         string              `bson:"uri"`             // 请求URI
}
