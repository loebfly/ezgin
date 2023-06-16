package redis

import "fmt"

// Yml 配置
type Yml struct {
	EZGinRedis EZGinRedis `koanf:"ezgin_redis"`
}

// SetYml 集合配置
type SetYml struct {
	RedisSet map[string]EZGinRedis `koanf:"ezgin_redis_set"`
}

type EZGinRedis struct {
	Addrs      []string `koanf:"addrs"`       // 连接地址, 优先级高于Host和Port, 选填，addrs和host/port二选一, 多个地址为集群模式
	Host       string   `koanf:"host"`        // 连接Ip, addrs为空时必填
	Port       int      `koanf:"port"`        // 连接端口, addrs为空时必填
	Password   string   `koanf:"password"`    // 密码, 必填
	Database   int      `koanf:"database"`    // 数据编号, 必填
	MasterName string   `koanf:"master_name"` // 哨兵模式主节点名称, 当为哨兵模式时必填
	Timeout    int      `koanf:"timeout"`     // 连接超时时间 默认5分钟
	Pool       struct {
		Min     int `koanf:"min"`     // 连接池最小连接数 默认3
		Max     int `koanf:"max"`     // 连接池最大连接数 默认20
		Idle    int `koanf:"idle"`    // 连接池最大空闲连接数 默认10
		Timeout int `koanf:"timeout"` // 连接池超时时间 默认60秒
	} `koanf:"pool" json:"pool"` // 连接池配置
	Tag string `koanf:"tag"` // 用于区分不同的数据库, 必填
}

func (receiver EZGinRedis) GetAddrs() []string {
	if len(receiver.Addrs) > 0 {
		return receiver.Addrs
	}
	return []string{fmt.Sprintf("%s:%d", receiver.Host, receiver.Port)}
}

func (receiver EZGinRedis) GetDB() int {
	if receiver.Database == -1 || len(receiver.Addrs) > 1 {
		return 0
	}
	return receiver.Database
}
