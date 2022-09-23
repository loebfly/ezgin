package nacos

import "errors"

var config = new(ymlConfig)

type ymlConfig struct {
	Nacos Yml `yaml:"nacos"`
}

func (cfg *ymlConfig) InitObj(obj Yml) error {
	if obj.Server == "" {
		return errors.New("nacos.server not null")
	}
	if obj.Port == "" {
		return errors.New("nacos.port not null")
	}

	if obj.LanNet == "" {
		return errors.New("nacos.lanNet not null")
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
	return nil
}
