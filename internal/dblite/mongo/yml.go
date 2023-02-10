package mongo

// Yml 配置
type Yml struct {
	EZGinMongo EZGinMongo `koanf:"ezgin_mongo"`
}

// SetYml 集合配置
type SetYml struct {
	MongoSet map[string]EZGinMongo `koanf:"ezgin_mongo_set"`
}

type EZGinMongo struct {
	Url      string `koanf:"url"`      // 连接地址, 必填
	Database string `koanf:"database"` // 数据库, 必填
	PoolMax  int    `koanf:"pool_max"` // 连接池最大连接数 默认20
	Tag      string `koanf:"tag"`      // 唯一标识，用于获取连接时查找使用, 必填
}
