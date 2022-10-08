package dblite

import (
	"github.com/go-redis/redis"
	mongoDB "github.com/loebfly/ezgin/internal/dblite/mongo"
	mysqlDB "github.com/loebfly/ezgin/internal/dblite/mysql"
	redisDB "github.com/loebfly/ezgin/internal/dblite/redis"
	"gopkg.in/mgo.v2"
	"gorm.io/gorm"
)

type enter int

const Enter = enter(0)

func InitDB(mongoObjs []mongoDB.EZGinMongo, mysqlObjs []mysqlDB.EZGinMysql, redisObjs []redisDB.EZGinRedis) {
	if mongoObjs != nil && len(mongoObjs) > 0 {
		mongoDB.InitObjs(mongoObjs)
	}
	if mysqlObjs != nil && len(mysqlObjs) > 0 {
		mysqlDB.InitObjs(mysqlObjs)
	}
	if redisObjs != nil && len(redisObjs) > 0 {
		redisDB.InitObjs(redisObjs)
	}
}

// IsExistMongoTag 判断是否存在mongo标签
func IsExistMongoTag(tag string) bool {
	return mongoDB.IsExistTag(tag)
}

func (enter) Mysql(tag ...string) (db *gorm.DB, err error) {
	return mysqlDB.GetDB(tag...)
}

func (enter) Mongo(tag ...string) (db *mgo.Database, returnDB func(db *mgo.Database), err error) {
	return mongoDB.GetDB(tag...)
}

func (enter) Redis(tag ...string) (db *redis.Client, err error) {
	return redisDB.GetDB(tag...)
}

func (enter) SafeExit() {
	mongoDB.Disconnect()
	mysqlDB.Disconnect()
	redisDB.Disconnect()
}
