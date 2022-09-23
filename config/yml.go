package config

import dbYml "github.com/loebfly/dblite/yml"

type GoYml struct {
	App struct {
		Name    string `yaml:"name"`     // 应用名称
		Port    int    `yaml:"port"`     // http服务端口
		PortSsl int    `yaml:"port_ssl"` // https服务端口
		Cert    string `yaml:"cert"`     // 证书, 用于https, 如果不需要https,则不需要配置
		Key     string `yaml:"key"`      // 私钥,用于https,如果不需要https,则不需要配置
		Debug   bool   `yaml:"debug"`    // 是否开启debug模式, 默认false, 如果开启, 则不会被其他服务调用
	} `yaml:"app"` // 应用配置

	Nacos struct {
		Server string `yaml:"server"` // nacos服务地址
		Env    string `yaml:"env"`    // 环境 test, dev, prod
		Yml    struct {
			Mysql string `yaml:"mysql"` // mysql配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
			Mongo string `yaml:"mongo"` // mongo配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
			Redis string `yaml:"redis"` // redis配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
		} `yaml:"yml"` // nacos配置文件名
	} `yaml:"nacos"` // nacos配置

	Log struct {
		Out   string `yaml:"out"`  // 日志输出方式, 可选值: console, file, mongo
		File  string `yaml:"file"` // 日志文件路径, 如果Out包含file, 则必须配置, 否则无法输出到文件
		Mongo struct {
			NacosMongo string      `yaml:"nacos_mongo"` // nacos的mongo日志配置文件名, 如果Out包含mongo, 则必须配置, 否则无法输出到mongo
			LocalMongo dbYml.Mongo `yaml:"local_mongo"` // 本地 mongo 数据库配置, 如果Out包含mongo, 则必须配置, 否则无法输出到mongo
		} `yaml:"mongo"` // mongo日志配置, nacos 与 local 二选一
	} `yaml:"log"` // 日志配置

	LocalDB struct {
		MySql dbYml.Mysql `yaml:"mysql"` // mysql 数据库本地配置
		Mongo dbYml.Mongo `yaml:"mongo"` // mongo 数据库本地配置
		Redis dbYml.Redis `yaml:"redis"` // redis 数据库本地配置
	} `yaml:"local_db"` // 本地数据库配置
}