package nacos

import "errors"

type Yml struct {
	Server      string `json:"server" yaml:"server"`             // 服务器ip地址 必填 多个用,分隔
	Port        string `json:"port" yaml:"port"`                 // 服务器端口 必填 多个用,分隔
	ClusterName string `json:"cluster_name" yaml:"cluster_name"` // 集群名称 默认为DEFAULT
	GroupName   string `json:"group_name" yaml:"group_name"`     // 组名称 默认为DEFAULT_GROUP
	Weight      int    `yaml:"weight"`                           // 权重 默认1
	Lan         bool   `yaml:"lan"`                              // 是否为内网 true:内网  false:外网
	LanNet      string `yaml:"lanNet"`                           // 网段前缀 必填 例如：192.168.3.
	App         AppYml `yaml:"app"`                              // 应用配置
}

type AppYml struct {
	Name    string `yaml:"name"`     // 应用名称
	Ip      string `yaml:"ip"`       // 应用ip地址, 默认为本机ip
	Port    int    `yaml:"port"`     // http服务端口
	PortSsl int    `yaml:"port_ssl"` // https服务端口
	Debug   bool   `yaml:"debug"`    // 是否开启debug模式, 默认false, 如果开启, 则不会被其他服务调用
}

type ymlConfig struct {
	Nacos Yml `yaml:"nacos"`
}

var config = new(ymlConfig)

func (cfg *ymlConfig) InitObj(obj Yml) error {
	if obj.Server == "" {
		return errors.New("nacos.server not null")
	}
	if obj.Port == "" {
		return errors.New("nacos.port not null")
	}

	if obj.LanNet == "" {
		return errors.New("nacos.lanNet not null")
	}

	if obj.ClusterName == "" {
		obj.ClusterName = "DEFAULT"
	}

	if obj.GroupName == "" {
		obj.GroupName = "DEFAULT_GROUP"
	}

	if obj.Weight == 0 {
		obj.Weight = 1
	}

	cfg.Nacos = obj
	return nil
}
