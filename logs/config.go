package logs

type Yml struct {
	// 日志输出方式 console,file
	Out string `yml:"out"`
	// 日志输出文件路径
	File string `yml:"file"`
}

type ymlConfig struct {
	Logs Yml `yml:"logs"`
}

var config = new(ymlConfig)

func (cfg *ymlConfig) initObj(obj Yml) error {
	cfg.Logs = obj
	cfg.fillNull()
	return nil
}

func (cfg *ymlConfig) fillNull() {
	if cfg.Logs.Out == "" {
		cfg.Logs.Out = OutConsole
	}
}
