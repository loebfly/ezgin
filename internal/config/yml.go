package config

import (
	"os"
	"path/filepath"
)

type Yml struct {
	EZGin EZGinYml `koanf:"ezgin"`
}

type EZGinYml struct {
	App struct {
		Name    string `koanf:"name"`     // 应用名称
		Ip      string `koanf:"ip"`       // 应用ip地址, 默认为本机ip
		Port    int    `koanf:"port"`     // http服务端口
		PortSsl int    `koanf:"port_ssl"` // https服务端口
		Cert    string `koanf:"cert"`     // 应用证书文件路径, 用于https, 如果不需要https,则不需要配置
		Key     string `koanf:"key"`      // 应用私钥文件路径, 用于https,如果不需要https,则不需要配置
		Debug   bool   `koanf:"debug"`    // 是否开启debug模式, 默认false, 如果开启, 则不会被其他服务调用
		Version string `koanf:"version"`  // 版本号
		Env     string `koanf:"env"`      // 环境 test, dev, prod
	} `koanf:"app"` // 应用配置

	Nacos struct {
		Server string `koanf:"server"` // nacos服务地址
		Yml    struct {
			// yml支持本地&nacos配置, 如果本地配置文件存在，则使用本地配置文件, 否则从nacos中读取
			// yml只需要配置文件的前缀，内部会自动拼接-$Env.yml, 例如: nacos-ezgin -> nacos-ezgin-test.yml
			// 配置模版参考: ezgin/template/xxx.yml
			Nacos string `koanf:"nacos"` // nacos配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml, 只支持单个配置文件, 如果不需要nacos配置文件,则不需要配置

			// 本地配置多个数据库，即数据库配置分开配置
			Mysql string `koanf:"mysql"` // mysql配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml, 多个配置文件用逗号分隔, 如果不需要mysql配置文件,则不需要配置
			Mongo string `koanf:"mongo"` // mongo配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml, 多个配置文件用逗号分隔, 如果不需要mongo配置文件,则不需要配置
			Redis string `koanf:"redis"` // redis配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml, 多个配置文件用逗号分隔, 如果不需要redis配置文件,则不需要配置
			Kafka string `koanf:"kafka"` // kafka配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml, 只支持单个配置文件, 如果不需要kafka配置文件,则不需要配置

			// 一个配置文件中配置多个数据库，即数据库配置合并到一个配置文件中
			MysqlSet string `koanf:"mysql_set"` // Mysql 集合配置文件名, 只需要配置文件的前缀，内部会自动拼接-$Env.yml
			MongoSet string `koanf:"mongo_set"` // Mongo 集合配置文件名, 只需要配置文件的前缀，内部会自动拼接-$Env.yml
			RedisSet string `koanf:"redis_set"` // Redis 集合配置文件名, 只需要配置文件的前缀，内部会自动拼接-$Env.yml
			KafkaSet string `koanf:"kafka_set"` // Kafka 集合配置文件名, 只需要配置文件的前缀，内部会自动拼接-$Env.yml
		} `koanf:"yml"` // nacos配置文件名
	} `koanf:"nacos"` // nacos配置

	Gin struct {
		Mode       string `koanf:"mode"`       // gin模式 debug, release
		Middleware string `koanf:"middleware"` // gin中间件, 用逗号分隔, 暂时支持 cors,trace,logs,recover 不填则默认全部开启, - 表示不开启
		MwLogs     struct {
			// 需要与Nacos.Yml.Mongo中配置文件名对应, 默认为Nacos.Yml.Mongo中第一个配置文件, - 表示不开启
			// 可配置变量: {header:xxx}，表示从请求头中获取key为xxx对应的值替换变量
			// 假设请求头中有一个key为: X-Request-Id，值为: abc123
			// 配置文件中的变量为: {header:X-Request-Id}
			// 则最终的变量值为: abc123
			MongoTag string `koanf:"mongo_tag"`
			// 日志表名, 默认为${App.Name}APIRequestLogs,
			// 可配置变量: {header:xxx}，表示从请求头中获取key为xxx对应的值替换变量
			// 假设请求头中有一个key为: X-Request-Id，值为: abc123
			// 效果为:{header:X-Request-Id}_APIRequestLogs -> abc123_APIRequestLogs
			MongoTable string `koanf:"mongo_table"`
			// kafka 消息主题, 默认为${App.Name}, 多个主题用逗号分隔, - 表示不开启
			// 可配置变量: {header:xxx}，表示从请求头中获取key为xxx对应的值替换变量
			// 假设请求头中有一个key为: X-Request-Id，值为: abc123
			// 效果为:{header:X-Request-Id}_Topic -> abc123_Topic
			KafkaTopic string `koanf:"kafka_topic"`
		} `koanf:"mw_logs"` // 日志中间件数据库配置
	} `yaml:"gin"` // gin配置

	Logs struct {
		Level string `koanf:"level"` // 日志级别 debug > info > warn > error, 默认debug即全部打印, - 表示不开启
		Out   string `koanf:"out"`   // 日志输出方式, 可选值: console, file 默认 console
		File  string `koanf:"file"`  // 日志文件路径, 如果Out包含file, 不填默认/opt/logs/${App.Name}, 生成的文件会带上.$(Date +%F).log
	} `koanf:"logs"` // 日志配置

	I18n struct {
		AppName    string `koanf:"app_name"`    // i18n应用名称, 多个用逗号分隔, 默认为default,${App.Name}, - 表示不开启
		ServerName string `koanf:"server_name"` // i18n微服务名称, 默认x-lang
		CheckUri   string `koanf:"check_uri"`   // i18n服务检查uri, 默认/lang/string/app/version
		QueryUri   string `koanf:"query_uri"`   // i18n服务查询uri, 默认/lang/string/list
		Duration   int    `koanf:"duration"`    //  i18n服务查询间隔, 默认360s
	} `koanf:"i18n"` // i18n配置
}

// GetYmlUrlOrPath 根据配置文件前缀获取nacos配置文件完整地址
func (yml EZGinYml) GetYmlUrlOrPath(prefix string) string {
	ymlName := prefix + "-" + yml.App.Env + ".yml"
	// 判断程序同级目录是否存在配置文件
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	ymlPath := filepath.Join(path, ymlName)
	if _, err := os.Stat(ymlPath); err == nil {
		return ymlPath
	}
	return yml.Nacos.Server + "nacos/v1/cs/configs?group=DEFAULT_GROUP&dataId=" + ymlName
}
