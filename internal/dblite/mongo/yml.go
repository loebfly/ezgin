package mongo

type Yml struct {
	Url      string `koanf:"url"`       // 连接地址, 必填
	Database string `koanf:"database"`  // 数据库, 必填
	PoolMax  int    `koanf:"pool_max"`  // 连接池最大连接数 默认20
	FindName string `koanf:"find_name"` // 用于区分不同的数据库, 必填
}
