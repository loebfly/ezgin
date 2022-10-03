package internal

import (
	"errors"
	"github.com/loebfly/ezgin/internal/nacos"
)

type ymlConfig struct {
	Nacos nacos.Yml `yaml:"nacos"`
}

var Config = new(ymlConfig)

func (cfg *ymlConfig) initObj(obj nacos.Yml) {
	if obj.Server == "" {
		panic(errors.New("nacos server is empty"))
	}
	if obj.Port == "" {
		panic(errors.New("nacos port is empty"))
	}

	if obj.LanNet == "" {
		panic(errors.New("nacos lanNet is empty"))
	}

	if obj.ClusterName == "" {
		obj.ClusterName = "DEFAULT"
	}

	if obj.GroupName == "" {
		obj.GroupName = "DEFAULT_GROUP"
	}

	if obj.Weight == 0 {
		obj.Weight = 1
	}

	cfg.Nacos = obj
}
