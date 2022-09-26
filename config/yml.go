package config

import (
	dbYml "github.com/loebfly/dblite/yml"
)

type Yml struct {
	EZGin struct {
		App struct {
			Name    string `yaml:"name"`     // 应用名称
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
				Mongo struct {
					log  string `yaml:"log"`  // 用于存储日志的mongo配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
					data string `yaml:"data"` // 用于存储数据的mongo配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
				} `yaml:"mongo"`
				Redis string `yaml:"redis"` // redis配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
			} `yaml:"yml"` // nacos配置文件名
		} `yaml:"nacos"` // nacos配置

		Gin struct {
			Mode       string `yaml:"mode"`       // gin模式 debug, release
			Middleware string `yaml:"middleware"` // 中间件, 用逗号分隔, 暂时支持cors, trace, logs
		} `yaml:"gin"` // gin配置

		Log struct {
			Out  string `yaml:"out"`  // 日志输出方式, 可选值: console, file 默认 console
			File string `yaml:"file"` // 日志文件路径, 如果Out包含file, 则必须配置, 否则无法输出到文件
		} `yaml:"log"` // 日志配置

		RequestLog struct {
			NacosMongo string      `yaml:"nacos_mongo"` // nacos的mongo日志配置文件名, 如果Out包含mongo, 则必须配置, 否则无法输出到mongo
			LocalMongo dbYml.Mongo `yaml:"local_mongo"` // 本地 mongo 数据库配置, 如果Out包含mongo, 则必须配置, 否则无法输出到mongo
		} `yaml:"request_log"` // 请求日志配置

		LocalDB struct {
			MySql dbYml.Mysql `yaml:"mysql"` // mysql 数据库本地配置
			Mongo dbYml.Mongo `yaml:"mongo"` // mongo 数据库本地配置
			Redis dbYml.Redis `yaml:"redis"` // redis 数据库本地配置
		} `yaml:"local_db"` // 本地数据库配置
	} `yaml:"ezgin"` // ezgin配置
}

func (yml Yml) GetNacosUrl(prefix string) string {
	return yml.EZGin.Nacos.Server + "nacos/v1/cs/configs?group=DEFAULT_GROUP&dataId=" + prefix + "-" + yml.EZGin.App.Env + ".yml"
}
