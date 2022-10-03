package config

import (
	"github.com/knadh/koanf"
	kYaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/internal/config/internal"
	"github.com/loebfly/ezgin/internal/logs"
	"gopkg.in/yaml.v3"
)

type enter int

const Enter = enter(0)

func InitPath(ymlPath string) {
	internal.InitPath(ymlPath)
}

func EZGin() EZGinYml {
	return internal.YmlObj.EZGin
}

// GetYmlData 以字节获取配置数据，结构体必须是yaml格式
func (enter) GetYmlData(confUrl string) (*koanf.Koanf, error) {
	resp, err := grequests.Get(confUrl, nil)
	if err != nil {
		logs.Enter.Error(confUrl + "配置下载失败! " + err.Error())
		return nil, err
	}
	conf := koanf.New(".")
	err = conf.Load(rawbytes.Provider([]byte(resp.String())), kYaml.Parser())
	if err != nil {
		logs.Enter.Error(confUrl + "配置格式解析错误:" + err.Error())
		return nil, err
	}
	return conf, nil
}

// GetYmlObj 以结构体获取配置数据，结构体tag必须包含json
func (enter) GetYmlObj(confUrl string, obj interface{}) error {
	resp, err := grequests.Get(confUrl, nil)
	if err != nil {
		logs.Enter.Error(confUrl + "配置下载失败! " + err.Error())
		return err
	}
	if resp.StatusCode != 200 {
		logs.Enter.Error(confUrl + "配置下载失败! " + resp.String())
		return err
	}
	err = yaml.Unmarshal(resp.Bytes(), &obj)
	if err != nil {
		logs.Enter.Error(confUrl + "配置格式解析错误:" + err.Error())
		return err
	}
	return nil
}

func (enter) GetString(key string) string {
	return internal.YmlData.String(key)
}

func (enter) GetInt(key string) int {
	return internal.YmlData.Int(key)
}

func (enter) GetBool(key string) bool {
	return internal.YmlData.Bool(key)
}

func (enter) GetFloat64(key string) float64 {
	return internal.YmlData.Float64(key)
}
