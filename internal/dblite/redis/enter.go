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
	} else {
		ezlogs.CDebug("REDIS", "初始化成功")
	}
	//ctl.addCheckTicker()
}

func GetDB(tag ...string) (db redis.UniversalClient, err error) {
	return ctl.getDB(tag...)
}

// NewPSub 创建一个新的PSub操作对象
func NewPSub(dbTag ...string) *PSubOperator {
	if len(dbTag) == 0 {
		dbTag = []string{}
	}
	return &PSubOperator{
		dbTag:    dbTag,
		channels: []string{},
	}
}

func GetAllTags() []string {
	var tags = make([]string, 0)
	for _, obj := range config.Objs {
		tags = append(tags, obj.Tag)
	}
	return tags
}

func Disconnect() {
	ctl.disconnect()
}
