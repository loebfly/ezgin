package redis

import (
	"errors"
	"fmt"
)

var config = new(ymlConfig)

type ymlConfig struct {
	Objs []Yml
}

func (cfg *ymlConfig) InitObjs(objs []Yml) {
	cfg.Objs = objs
	err := cfg.checkObjs()
	if err != nil {
		panic(err)
	}
	cfg.fillNull()
}

func (cfg *ymlConfig) checkObjs() error {
	for i, obj := range cfg.Objs {
		if obj.Host == "" {
			return errors.New(fmt.Sprintf("第 %d 个 redis.host 不可为空", i+1))
		}
		if obj.Port == 0 {
			return errors.New(fmt.Sprintf("第 %d 个 redis.port 不可为空", i+1))
		}
		if obj.Password == "" {
			return errors.New(fmt.Sprintf("第 %d 个 redis.password 不可为空", i+1))
		}
	}
	return nil
}

func (cfg *ymlConfig) fillNull() {
	for i, obj := range cfg.Objs {
		if obj.Timeout == 0 {
			cfg.Objs[i].Timeout = 1000
		}
		if obj.Pool.Min == 0 {
			cfg.Objs[i].Pool.Min = 3
		}
		if obj.Pool.Max == 0 {
			cfg.Objs[i].Pool.Max = 20
		}
		if obj.Pool.Idle == 0 {
			cfg.Objs[i].Pool.Idle = 10
		}
		if obj.Pool.Timeout == 0 {
			cfg.Objs[i].Pool.Timeout = 300
		}
	}
}
