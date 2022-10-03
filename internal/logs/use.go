package logs

var (
	WillOutputHandlers = make([]GetExtraInfo, 0)
)

// GetExtraInfo 当日志即将输出时触发的回调方法
// 返回的是一个map, 可以用来添加额外的信息, key为信息, value为第几个位置
type GetExtraInfo func(category, level string) map[string]int

func Use(handler GetExtraInfo) {
	WillOutputHandlers = append(WillOutputHandlers, handler)
}
