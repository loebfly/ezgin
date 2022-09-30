package config

import (
	"fmt"
	"github.com/knadh/koanf"
	kYaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/logs"
	"gopkg.in/yaml.v3"
)

type enter int

const Enter = enter(0)

var (
	YmlData *koanf.Koanf
	YmlObj  *Yml
)

func (enter) InitPath(ymlPath string) {
	logs.Enter.Debug("读取配置文件:" + ymlPath)
	YmlData = koanf.New(".")
	f := file.Provider(ymlPath)
	err := YmlData.Load(f, kYaml.Parser())
	if err != nil {
		panic(fmt.Errorf("配置文件解析错误:%s", err.Error()))
	}
	fBytes, err := f.ReadBytes()
	if err != nil {
		panic(fmt.Errorf("读取配置文件错误:%s", err.Error()))
	}
	if err = yaml.Unmarshal(fBytes, &YmlObj); err != nil {
		panic(fmt.Errorf("配置文件解析错误:%s", err.Error()))
	}
}

func (enter) checkYmlObj() {
	if YmlObj == nil {
		panic("未初始化配置文件")
	}
	if YmlObj.EZGin.App.Name == "" {
		panic("未初始化配置文件")
	}
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
