package logs

type Yml struct {
	// 日志输出方式 console,file
	Out string `yml:"out"`
	// 日志输出文件路径
	File string `yml:"file"`
}
