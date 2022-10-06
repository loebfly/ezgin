package mysql

type Yml struct {
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
	FindName string `koanf:"find_name"` // 用于区分不同的数据库, 必填
}
