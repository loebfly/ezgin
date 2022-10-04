package config

import (
	"fmt"
	"github.com/knadh/koanf"
	kYaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/loebfly/ezgin/internal/logs"
)

var (
	YmlData *koanf.Koanf
	YmlObj  *Yml
)

func InitPath(ymlPath string) {
	logs.Enter.Debug("读取配置文件:" + ymlPath)
	YmlData = koanf.New(".")
	f := file.Provider(ymlPath)
	err := YmlData.Load(f, kYaml.Parser())
	if err != nil {
		panic(fmt.Errorf("配置文件解析错误:%s", err.Error()))
	}
	if err = YmlData.Unmarshal("", &YmlObj); err != nil {
		panic(fmt.Errorf("配置文件解析错误:%s", err.Error()))
	}
	checkYmlObj()
}

func checkYmlObj() {
	if YmlObj == nil {
		panic("未初始化配置文件")
	}
	if YmlObj.EZGin.App.Name == "" {
		panic("未初始化配置文件")
	}
}
