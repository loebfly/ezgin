package nacos

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/levigross/grequests"
	"github.com/loebfly/klogs"
)

func (enter) GetConfig(confUrl string) (*koanf.Koanf, error) {
	resp, err := grequests.Get(confUrl, nil)
	if err != nil {
		klogs.Error("MySQL配置下载失败! " + err.Error())
		return nil, err
	}
	conf := koanf.New(".")
	err = conf.Load(rawbytes.Provider([]byte(resp.String())), yaml.Parser())
	if err != nil {
		klogs.Error("MySQL配置格式解析错误:" + err.Error())
		return nil, err
	}
	return conf, nil
}
