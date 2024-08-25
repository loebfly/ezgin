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

// GetYmlData 以字节请求yaml格式配置数据
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

// GetString 获取key对应的string类型值
func GetString(key string) string {
	return config.YmlData.String(key)
}

// GetInt 获取key对应的int类型值
func GetInt(key string) int {
	return config.YmlData.Int(key)
}

// GetInt64 获取key对应的int64类型的值
func GetInt64(key string) int64 {
	return config.YmlData.Int64(key)
}

// GetBool 获取key对应的bool类型值
func GetBool(key string) bool {
	return config.YmlData.Bool(key)
}

// GetFloat64 获取key对应的float64类型值
func GetFloat64(key string) float64 {
	return config.YmlData.Float64(key)
}

// GetArray 获取数组
func GetArray[T any](key string) []T {
	var res []T
	if srcArr, ok := config.YmlData.Get(key).([]interface{}); ok {
		res = make([]T, len(srcArr))
		for i := 0; i < len(srcArr); i++ {
			if val, valOk := srcArr[i].(T); valOk {
				res[i] = val
			} else {
				var place T
				res[i] = place
			}

		}
		return res
	}
	return res
}

// GetYmlUrlOrPath 获取前缀对应的配置文件路径或Nacos的URL
func GetYmlUrlOrPath(prefix string) string {
	return config.YmlObj.EZGin.GetYmlUrlOrPath(prefix)
}
