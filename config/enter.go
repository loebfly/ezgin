package config

import (
	"github.com/knadh/koanf"
	kYaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/levigross/grequests"
	"github.com/loebfly/klogs"
	"gopkg.in/yaml.v3"
)

type enter int

const Enter = enter(0)

var YmlData *koanf.Koanf

var YmlObj *Yml

func (enter) Init(ymlPath string) error {
	klogs.Debug("读取配置文件:" + ymlPath)
	YmlData = koanf.New(".")
	f := file.Provider(ymlPath)
	err := YmlData.Load(f, kYaml.Parser())
	if err != nil {
		klogs.Error("读取配置文件错误:" + err.Error())
		return err
	}
	fBytes, err := f.ReadBytes()
	if err != nil {
		klogs.Error("读取配置文件错误:" + err.Error())
		return err
	}
	if err = yaml.Unmarshal(fBytes, &YmlObj); err != nil {
		klogs.Error("配置文件解析错误:" + err.Error())
		return err
	}
	return nil
}

// GetYmlData 以字节获取配置数据，结构体必须是yaml格式
func (enter) GetYmlData(confUrl string) (*koanf.Koanf, error) {
	resp, err := grequests.Get(confUrl, nil)
	if err != nil {
		klogs.Error(confUrl + "配置下载失败! " + err.Error())
		return nil, err
	}
	conf := koanf.New(".")
	err = conf.Load(rawbytes.Provider([]byte(resp.String())), kYaml.Parser())
	if err != nil {
		klogs.Error(confUrl + "配置格式解析错误:" + err.Error())
		return nil, err
	}
	return conf, nil
}

// GetYmlObj 以结构体获取配置数据，结构体tag必须包含json
func (enter) GetYmlObj(confUrl string, obj interface{}) error {
	resp, err := grequests.Get(confUrl, nil)
	if err != nil {
		klogs.Error(confUrl + "配置下载失败! " + err.Error())
		return err
	}
	if resp.StatusCode != 200 {
		klogs.Error(confUrl + "配置下载失败! " + resp.String())
		return err
	}
	err = yaml.Unmarshal(resp.Bytes(), &obj)
	if err != nil {
		klogs.Error(confUrl + "配置格式解析错误:" + err.Error())
		return err
	}
	return nil
}
