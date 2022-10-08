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

func (enter) Mysql(findName ...string) (db *gorm.DB, err error) {
	return mysqlDB.GetDB(findName...)
}

func (enter) Mongo(findName ...string) (db *mgo.Database, returnDB func(db *mgo.Database), err error) {
	return mongoDB.GetDB(findName...)
}

func (enter) Redis(findName ...string) (db *redis.Client, err error) {
	return redisDB.GetDB(findName...)
}

func (enter) SafeExit() {
	mongoDB.Disconnect()
	mysqlDB.Disconnect()
	redisDB.Disconnect()
}
