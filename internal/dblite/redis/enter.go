package redis

import (
	"github.com/go-redis/redis"
	"github.com/loebfly/ezgin/internal/logs"
)

func InitObjs(objs []EZGinRedis) {
	logs.Enter.CDebug("REDIS", "初始化")
	config.InitObjs(objs)
	err := ctl.initConnect()
	if err != nil {
		logs.Enter.CError("REDIS", "初始化失败: {}", err.Error())
	}
	logs.Enter.CDebug("REDIS", "初始化成功")
	ctl.addCheckTicker()
}

func GetDB(tag ...string) (db *redis.Client, err error) {
	return ctl.getDB(tag...)
}

func Disconnect() {
	ctl.disconnect()
}
