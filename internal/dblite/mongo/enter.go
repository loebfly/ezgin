package mongo

import (
	"github.com/loebfly/ezgin/internal/logs"
	"gopkg.in/mgo.v2"
)

func InitObjs(objs []Yml) {
	logs.Enter.CDebug("MONGO", "初始化")
	config.InitObjs(objs)
	err := ctl.initConnect()
	if err != nil {
		logs.Enter.CError("MONGO", "初始化失败: %s", err.Error())
	}
	ctl.addCheckTicker()
}

func GetDB(fineName ...string) (db *mgo.Database, returnDB func(db *mgo.Database), err error) {
	return ctl.getDB(fineName...)
}

func Disconnect() {
	ctl.disconnect()
}
