package internal

import "github.com/loebfly/ezgin/internal/logs"

type ymlConfig struct {
	Logs logs.Yml `yml:"logs"`
}

var Config = new(ymlConfig)

func (cfg *ymlConfig) initObj(obj logs.Yml) {
	cfg.Logs = obj
	cfg.fillNull()
}

func (cfg *ymlConfig) fillNull() {
	if cfg.Logs.Out == "" {
		cfg.Logs.Out = logs.OutConsole
	}
}
