package internal

import (
	"fmt"
	"github.com/knadh/koanf"
	kYaml "github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/loebfly/ezgin/internal/config"
	"github.com/loebfly/ezgin/internal/logs"
	"gopkg.in/yaml.v3"
)

var (
	YmlData *koanf.Koanf
	YmlObj  *config.Yml
)

func initPath(ymlPath string) {
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
