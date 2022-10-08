package mysql

import (
	"errors"
	"fmt"
)

var config = new(ymlConfig)

type ymlConfig struct {
	Objs []EZGinMysql
}

func (cfg *ymlConfig) InitObjs(objs []EZGinMysql) {
	cfg.Objs = objs
	err := cfg.checkObjs()
	if err != nil {
		panic(err)
	}
	cfg.fillNull()
}

func (cfg *ymlConfig) checkObjs() error {
	for i, obj := range cfg.Objs {
		if obj.Url == "" {
			return errors.New(fmt.Sprintf("第 %d 个 mysql url 不可为空", i+1))
		}
	}
	return nil
}

func (cfg *ymlConfig) fillNull() {
	for i, obj := range cfg.Objs {
		if obj.Pool.Max == 0 {
			cfg.Objs[i].Pool.Max = 20
		}
		if obj.Pool.Idle == 0 {
			cfg.Objs[i].Pool.Idle = 10
		}
		if obj.Pool.Timeout.Life == 0 {
			cfg.Objs[i].Pool.Timeout.Life = 60
		}
		if obj.Pool.Timeout.Idle == 0 {
			cfg.Objs[i].Pool.Timeout.Idle = 60
		}
	}
}
