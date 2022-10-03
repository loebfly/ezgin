package dblite

type MongoYml struct {
	Url      string `yaml:"url" json:"url"`
	Database string `yaml:"database" json:"database"`
	PoolMax  int    `yaml:"pool_max" json:"pool_max"`
}

type MysqlYml struct {
	Url   string `yaml:"url" json:"url"`
	Debug bool   `yaml:"debug" json:"debug"`
	Pool  struct {
		Max     int `yaml:"max" json:"max"`
		Idle    int `yaml:"idle" json:"idle"`
		Timeout struct {
			Idle int `yaml:"idle" json:"idle"`
			Life int `yaml:"life" json:"life"`
		} `yaml:"timeout" json:"timeout"`
	} `yaml:"pool" json:"pool"`
}

type RedisYml struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	Database int    `yaml:"database" json:"database"`
	Timeout  int    `yaml:"timeout" json:"timeout"`
	Pool     struct {
		Min     int `yaml:"min" json:"min"`
		Max     int `yaml:"max" json:"max"`
		Idle    int `yaml:"idle" json:"idle"`
		Timeout int `yaml:"timeout" json:"timeout"`
	} `yaml:"pool" json:"pool"`
}
