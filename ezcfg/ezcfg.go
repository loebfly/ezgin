package ezcfg

import (
	"github.com/knadh/koanf"
	kYaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/levigross/grequests"
	"github.com/loebfly/ezgin/ezlogs"
	"github.com/loebfly/ezgin/internal/config"
	"strings"
)

// GetYmlData 以字节获取配置数据，结构体必须是yaml格式
func GetYmlData(confUrl string) (*koanf.Koanf, error) {
	resp, err := grequests.Get(confUrl, nil)
	if err != nil {
		ezlogs.Error(confUrl + "配置下载失败! " + err.Error())
		return nil, err
	}
	conf := koanf.New(".")
	err = conf.Load(rawbytes.Provider([]byte(resp.String())), kYaml.Parser())
	if err != nil {
		ezlogs.Error(confUrl + "配置格式解析错误:" + err.Error())
		return nil, err
	}
	return conf, nil
}

// GetYmlObj 以结构体获取配置数据，结构体tag必须包含json
func GetYmlObj(confUrlOrPath string, obj any) error {
	// 判断confUrl是否是本地路径
	if strings.HasPrefix(confUrlOrPath, "http") {
		resp, err := grequests.Get(confUrlOrPath, nil)
		if err != nil {
			ezlogs.Error(confUrlOrPath + "配置下载失败! " + err.Error())
			return err
		}
		if resp.StatusCode != 200 {
			ezlogs.Error(confUrlOrPath + "配置下载失败! " + resp.String())
			return err
		}
		conf := koanf.New(".")
		err = conf.Load(rawbytes.Provider([]byte(resp.String())), kYaml.Parser())
		if err != nil {
			ezlogs.Error(confUrlOrPath + "配置格式解析错误:" + err.Error())
			return err
		}
		err = conf.Unmarshal("", &obj)
		if err != nil {
			ezlogs.Error(confUrlOrPath + "配置格式解析错误:" + err.Error())
			return err
		}
	} else {
		conf := koanf.New(".")
		f := file.Provider(confUrlOrPath)
		err := conf.Load(f, kYaml.Parser())
		if err != nil {
			ezlogs.Error(confUrlOrPath + "配置格式解析错误:" + err.Error())
			return err
		}
		err = conf.Unmarshal("", &obj)
		if err != nil {
			ezlogs.Error(confUrlOrPath + "配置格式解析错误:" + err.Error())
			return err
		}
	}

	return nil
}

func GetString(key string) string {
	return config.YmlData.String(key)
}

func GetInt(key string) int {
	return config.YmlData.Int(key)
}

func GetBool(key string) bool {
	return config.YmlData.Bool(key)
}

func GetFloat64(key string) float64 {
	return config.YmlData.Float64(key)
}
