package mysql

import (
	"github.com/loebfly/ezgin/ezlogs"
	"gorm.io/gorm"
)

func InitObjs(objs []EZGinMysql) {
	ezlogs.CDebug("MYSQL", "初始化")
	config.InitObjs(objs)
	err := ctl.initConnect()
	if err != nil {
		ezlogs.CError("MYSQL", "初始化失败: {}", err.Error())
	}
	ezlogs.CInfo("MYSQL", "初始化成功")
	ctl.addCheckTicker()
}

func GetDB(tag ...string) (db *gorm.DB, err error) {
	return ctl.getDB(tag...)
}

func Disconnect() {
	ctl.disconnect()
}
