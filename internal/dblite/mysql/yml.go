package mysql

// Yml 配置
type Yml struct {
	EZGinMysql EZGinMysql `koanf:"ezgin_mysql"`
}

// SetYml 集合配置
type SetYml struct {
	MysqlSet map[string]EZGinMysql `koanf:"ezgin_mysql_set"`
}

type EZGinMysql struct {
	Url   string `koanf:"url"`   // 连接地址, 必填
	Debug bool   `koanf:"debug"` // 是否开启调试模式，默认false
	Pool  struct {
		Max     int `koanf:"max"`  // 连接池最大连接数 默认20
		Idle    int `koanf:"idle"` // 连接池最大空闲连接数 默认10
		Timeout struct {
			Idle int `koanf:"idle"` // 连接池最大空闲时间 默认60s
			Life int `koanf:"life"` // 连接池最大生存时间 默认60s
		} `koanf:"timeout"` // 连接池超时时间
	} `koanf:"pool"` // 连接池配置
	Tag string `koanf:"tag"` // 唯一标识，用于获取连接时查找使用, 必填
}
