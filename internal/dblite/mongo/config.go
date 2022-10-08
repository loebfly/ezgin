package mongo

import (
	"errors"
	"fmt"
)

var config = new(ymlConfig)

type ymlConfig struct {
	Objs []EZGinMongo
}

func (cfg *ymlConfig) InitObjs(objs []EZGinMongo) {
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
			return errors.New(fmt.Sprintf("第 %d 个 mongo.url 不可为空", i+1))
		}
		if obj.Database == "" {
			return errors.New(fmt.Sprintf("第 %d 个 mongo.database 不可为空", i+1))
		}
	}
	return nil
}

func (cfg *ymlConfig) fillNull() {
	for i, obj := range cfg.Objs {
		if obj.PoolMax == 0 {
			cfg.Objs[i].PoolMax = 20
		}
	}
}

func (cfg *ymlConfig) isExistTag(tag string) bool {
	for _, obj := range cfg.Objs {
		if obj.Tag == tag {
			return true
		}
	}
	return false
}
