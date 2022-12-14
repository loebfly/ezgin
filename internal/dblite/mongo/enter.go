package mongo

import (
	"github.com/loebfly/ezgin/internal/logs"
	"gopkg.in/mgo.v2"
)

func InitObjs(objs []EZGinMongo) {
	logs.Enter.CDebug("MONGO", "初始化")
	config.InitObjs(objs)
	err := ctl.initConnect()
	if err != nil {
		logs.Enter.CError("MONGO", "初始化失败: {}", err.Error())
	}
	logs.Enter.CInfo("MONGO", "初始化成功")
	ctl.addCheckTicker()
}

func GetDB(tag ...string) (db *mgo.Database, returnDB func(db *mgo.Database), err error) {
	return ctl.getDB(tag...)
}

func Disconnect() {
	ctl.disconnect()
}
