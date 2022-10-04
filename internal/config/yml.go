package config

type Yml struct {
	EZGin EZGinYml `koanf:"ezgin"`
}

type EZGinYml struct {
	App struct {
		Name    string `koanf:"name"`     // 应用名称
		Ip      string `koanf:"ip"`       // 应用ip地址, 默认为本机ip
		Port    int    `koanf:"port"`     // http服务端口
		PortSsl int    `koanf:"port_ssl"` // https服务端口
		Cert    string `koanf:"cert"`     // 证书, 用于https, 如果不需要https,则不需要配置
		Key     string `koanf:"key"`      // 私钥,用于https,如果不需要https,则不需要配置
		Debug   bool   `koanf:"debug"`    // 是否开启debug模式, 默认false, 如果开启, 则不会被其他服务调用
		Version string `koanf:"version"`  // 版本号
		Env     string `koanf:"env"`      // 环境 test, dev, prod
	} `koanf:"app"` // 应用配置

	Nacos struct {
		Server string `koanf:"server"` // nacos服务地址
		Yml    struct {
			Nacos string `koanf:"nacos"` // nacos配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
			Mysql string `koanf:"mysql"` // mysql配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
			Mongo string `koanf:"mongo"` // mongo配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
			Redis string `koanf:"redis"` // redis配置文件名 只需要配置文件的前缀，内部会自动拼接-$Env.yml
		} `koanf:"yml"` // nacos配置文件名
	} `koanf:"nacos"` // nacos配置

	Gin struct {
		Mode       string `koanf:"mode"`       // gin模式 debug, release
		Middleware string `koanf:"middleware"` // 中间件, 用逗号分隔, 暂时支持cors, trace, logs 不填则默认全部开启, - 表示不开启
	} `yaml:"gin"` // gin配置

	Logs struct {
		Out  string `koanf:"out"`  // 日志输出方式, 可选值: console, file 默认 console
		File string `koanf:"file"` // 日志文件路径, 如果Out包含file, 不填默认/opt/logs/${App.Name}.$(Date +%F).log
	} `yaml:"log"` // 日志配置

	//LocalDB struct {
	//	MySql dbYml.Mysql `koanf:"mysql"` // mysql 数据库本地配置
	//	Mongo dbYml.Mongo `koanf:"mongo"` // mongo 数据库本地配置
	//	Redis dbYml.Redis `koanf:"redis"` // redis 数据库本地配置
	//} `yaml:"local_db"` // 本地数据库配置
}

// GetNacosUrl 根据配置文件前缀获取nacos配置文件完整地址
func (yml EZGinYml) GetNacosUrl(prefix string) string {
	return yml.Nacos.Server + "nacos/v1/cs/configs?group=DEFAULT_GROUP&dataId=" + prefix + "-" + yml.App.Env + ".yml"
}
