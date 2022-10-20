package logs

type Yml struct {
	Level string
	// 日志输出方式 console,file
	Out string
	// 日志输出文件路径
	File string
}
