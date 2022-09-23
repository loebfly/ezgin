package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/loebfly/klogs"
)

var cnf *koanf.Koanf

func (enter) Load(ymlPath string) error {
	klogs.Debug("读取配置文件:" + ymlPath)
	cnf = koanf.New(".")
	f := file.Provider(ymlPath)
	err := cnf.Load(f, yaml.Parser())
	if err != nil {
		klogs.Error("读取配置文件错误:" + err.Error())
		return err
	}
	return nil
}
