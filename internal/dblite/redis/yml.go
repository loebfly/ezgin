package redis

type Yml struct {
	EZGinRedis EZGinRedis `koanf:"ezgin_redis"`
}

type EZGinRedis struct {
	Host     string `koanf:"host"`     // 连接Ip, 必填
	Port     int    `koanf:"port"`     // 连接端口, 必填
	Password string `koanf:"password"` // 密码, 必填
	Database int    `koanf:"database"` // 数据编号, 必填
	Timeout  int    `koanf:"timeout"`  // 连接超时时间 默认1000ms
	Pool     struct {
		Min     int `koanf:"min"`     // 连接池最小连接数 默认3
		Max     int `koanf:"max"`     // 连接池最大连接数 默认20
		Idle    int `koanf:"idle"`    // 连接池最大空闲连接数 默认10
		Timeout int `koanf:"timeout"` // 连接池超时时间 默认300ms
	} `koanf:"pool" json:"pool"` // 连接池配置
	Tag string `koanf:"tag"` // 用于区分不同的数据库, 必填
}
