package mongo

import (
	"github.com/loebfly/ezgin/ezlogs"
	"gopkg.in/mgo.v2"
)

func InitObjs(objs []EZGinMongo) {
	ezlogs.CDebug("MONGO", "初始化")
	config.InitObjs(objs)
	err := ctl.initConnect()
	if err != nil {
		ezlogs.CError("MONGO", "初始化失败: {}", err.Error())
	}
	ezlogs.CInfo("MONGO", "初始化成功")
	ctl.addCheckTicker()
}

func GetDB(tag ...string) (db *mgo.Database, returnDB func(db *mgo.Database), err error) {
	return ctl.getDB(tag...)
}

func Disconnect() {
	ctl.disconnect()
}
