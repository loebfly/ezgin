package config

import (
	"github.com/knadh/koanf"
	kYaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/internal/logs"
	"strings"
)

type enter int

const Enter = enter(0)

func EZGin() EZGinYml {
	return YmlObj.EZGin
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
func (enter) GetYmlObj(confUrlOrPath string, obj any) error {
	// 判断confUrl是否是本地路径
	if strings.HasPrefix(confUrlOrPath, "http") {
		resp, err := grequests.Get(confUrlOrPath, nil)
		if err != nil {
			logs.Enter.Error(confUrlOrPath + "配置下载失败! " + err.Error())
			return err
		}
		if resp.StatusCode != 200 {
			logs.Enter.Error(confUrlOrPath + "配置下载失败! " + resp.String())
			return err
		}
		conf := koanf.New(".")
		err = conf.Load(rawbytes.Provider([]byte(resp.String())), kYaml.Parser())
		if err != nil {
			logs.Enter.Error(confUrlOrPath + "配置格式解析错误:" + err.Error())
			return err
		}
		err = conf.Unmarshal("", &obj)
		if err != nil {
			logs.Enter.Error(confUrlOrPath + "配置格式解析错误:" + err.Error())
			return err
		}
	} else {
		conf := koanf.New(".")
		f := file.Provider(confUrlOrPath)
		err := conf.Load(f, kYaml.Parser())
		if err != nil {
			logs.Enter.Error(confUrlOrPath + "配置格式解析错误:" + err.Error())
			return err
		}
		err = conf.Unmarshal("", &obj)
		if err != nil {
			logs.Enter.Error(confUrlOrPath + "配置格式解析错误:" + err.Error())
			return err
		}
	}

	return nil
}

func (enter) GetString(key string) string {
	return YmlData.String(key)
}

func (enter) GetInt(key string) int {
	return YmlData.Int(key)
}

func (enter) GetBool(key string) bool {
	return YmlData.Bool(key)
}

func (enter) GetFloat64(key string) float64 {
	return YmlData.Float64(key)
}
