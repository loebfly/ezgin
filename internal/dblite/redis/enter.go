package redis

import (
	"github.com/go-redis/redis"
	"github.com/loebfly/ezgin/ezlogs"
)

func InitObjs(objs []EZGinRedis) {
	ezlogs.CDebug("REDIS", "初始化")
	config.InitObjs(objs)
	err := ctl.initConnect()
	if err != nil {
		ezlogs.CError("REDIS", "初始化失败: {}", err.Error())
	}
	ezlogs.CDebug("REDIS", "初始化成功")
	ctl.addCheckTicker()
}

func GetDB(tag ...string) (db *redis.Client, err error) {
	return ctl.getDB(tag...)
}

func Disconnect() {
	ctl.disconnect()
}
