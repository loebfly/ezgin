package mysql

import (
	"github.com/loebfly/ezgin/internal/logs"
	"gorm.io/gorm"
)

func InitObjs(objs []EZGinMysql) {
	logs.Enter.CDebug("MYSQL", "初始化")
	config.InitObjs(objs)
	err := ctl.initConnect()
	if err != nil {
		logs.Enter.CError("MYSQL", "初始化失败: {}", err.Error())
	}
	logs.Enter.CInfo("MYSQL", "初始化成功")
	ctl.addCheckTicker()
}

func GetDB(tag ...string) (db *gorm.DB, err error) {
	return ctl.getDB(tag...)
}

func Disconnect() {
	ctl.disconnect()
}
