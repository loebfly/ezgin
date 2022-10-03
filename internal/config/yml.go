package config

import dbYml "github.com/loebfly/dblite/yml"

type Yml struct {
	EZGin EZGinYml `yaml:"ezgin"`
}

type EZGinYml struct {
	App struct {
		Name    string `yaml:"name"`     // 应用名称
		Ip      string `yaml:"ip"`       // 应用ip地址, 默认为本机ip
		Port    int    `yaml:"port"`     // http服务端口
		PortSsl int    `yaml:"port_ssl"` // https服务端口
		Cert    string `yaml:"cert"`     // 证书, 用于https, 如果不需要https,则不需要配置
		Key     string `yaml:"key"`      // 私钥,用于https,如果不需要https,则不需要配置
		Debug   bool   `yaml:"debug"`    // 是否开启debug模式, 默认false, 如果开启, 则不会被其他服务调用
		Version string `yaml:"version"`  // 版本号
		Env     string `yaml:"env"`      // 环境 test, dev, prod
	} `yaml:"app"` // 应用配置

	Nacos struct {
		Server string `yaml:"server"` // nacos服务地址
		Yml    struct {
			Nacos string `yaml:"nacos"` // nacos配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
			Mysql string `yaml:"mysql"` // mysql配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
			Mongo string `yaml:"mongo"` // mongo配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
			Redis string `yaml:"redis"` // redis配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
		} `yaml:"yml"` // nacos配置文件名
	} `yaml:"nacos"` // nacos配置

	Gin struct {
		Mode       string `yaml:"mode"`       // gin模式 debug, release
		Middleware string `yaml:"middleware"` // 中间件, 用逗号分隔, 暂时支持cors, trace, logs 不填则默认全部开启, - 表示不开启
	} `yaml:"gin"` // gin配置

	Logs struct {
		Out  string `yaml:"out"`  // 日志输出方式, 可选值: console, file 默认 console
		File string `yaml:"file"` // 日志文件路径, 如果Out包含file, 不填默认/opt/logs/${App.Name}.$(Date +%F).log
	} `yaml:"log"` // 日志配置

	LocalDB struct {
		MySql dbYml.Mysql `yaml:"mysql"` // mysql 数据库本地配置
		Mongo dbYml.Mongo `yaml:"mongo"` // mongo 数据库本地配置
		Logs  dbYml.Mongo `yaml:"mongo"` // mongo 数据库本地配置, 用于日志输出
		Redis dbYml.Redis `yaml:"redis"` // redis 数据库本地配置
	} `yaml:"local_db"` // 本地数据库配置
}

// GetNacosUrl 根据配置文件前缀获取nacos配置文件完整地址
func (yml EZGinYml) GetNacosUrl(prefix string) string {
	return yml.Nacos.Server + "nacos/v1/cs/configs?group=DEFAULT_GROUP&dataId=" + prefix + "-" + yml.App.Env + ".yml"
}
