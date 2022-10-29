package kafka

import "errors"

var config = new(ymlConfig)

type ymlConfig struct {
	Obj EZGinKafka
}

func (cfg *ymlConfig) initObj(obj EZGinKafka) {
	cfg.Obj = obj
	err := cfg.checkObj()
	if err != nil {
		panic(err)
	}
	cfg.fillNull()
}

func (cfg *ymlConfig) checkObj() error {
	if cfg.Obj.Servers == "" {
		return errors.New("kafka.Servers 不可为空")
	}
	return nil
}

func (cfg *ymlConfig) fillNull() {
	if cfg.Obj.Ack == "" {
		cfg.Obj.Ack = "all"
	}
	if cfg.Obj.Partitioner == "" {
		cfg.Obj.Partitioner = "hash"
	}
	if cfg.Obj.Version == "" {
		cfg.Obj.Version = "2.8.1"
	}
}
