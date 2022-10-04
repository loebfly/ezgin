package logs

type ymlConfig struct {
	Logs Yml `yml:"logs"`
}

var Config = new(ymlConfig)

func (cfg *ymlConfig) initObj(obj Yml) {
	cfg.Logs = obj
	cfg.fillNull()
}

func (cfg *ymlConfig) fillNull() {
	if cfg.Logs.Out == "" {
		cfg.Logs.Out = OutConsole
	}
}
