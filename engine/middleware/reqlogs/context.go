package reqlogs

type ReqCtx struct {
	RequestId   string      // 请求id
	ReqTime     string      // 请求时间
	ReqParams   interface{} // 请求参数
	RespTime    string      // 响应时间
	RespParams  interface{} // 响应参数
	TTL         int         // 请求耗时
	Method      string      // 请求方法
	ContentType string      // 请求类型
	URI         string      // 请求URI
}
