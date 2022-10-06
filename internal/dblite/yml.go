package dblite

type MongoYml struct {
	Url      string `koanf:"url"`       // 连接地址, 必填
	Database string `koanf:"database"`  // 数据库, 必填
	PoolMax  int    `koanf:"pool_max"`  // 连接池最大连接数 默认20
	FindName string `koanf:"find_name"` // 用于区分不同的数据库, 必填
}

type MysqlYml struct {
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

type RedisYml struct {
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
	FindName string `koanf:"find_name"` // 用于区分不同的数据库, 必填
}
