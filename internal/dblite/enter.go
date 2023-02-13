package dblite

import (
	kafkaDB "github.com/loebfly/ezgin/internal/dblite/kafka"
	mongoDB "github.com/loebfly/ezgin/internal/dblite/mongo"
	mysqlDB "github.com/loebfly/ezgin/internal/dblite/mysql"
	redisDB "github.com/loebfly/ezgin/internal/dblite/redis"
)

// InitDB 初始化数据库
func InitDB(mongoObjs []mongoDB.EZGinMongo, mysqlObjs []mysqlDB.EZGinMysql, redisObjs []redisDB.EZGinRedis, kafkaObjs []kafkaDB.EZGinKafka) {
	if mongoObjs != nil && len(mongoObjs) > 0 {
		mongoDB.InitObjs(mongoObjs)
	}
	if mysqlObjs != nil && len(mysqlObjs) > 0 {
		mysqlDB.InitObjs(mysqlObjs)
	}
	if redisObjs != nil && len(redisObjs) > 0 {
		redisDB.InitObjs(redisObjs)
	}
	if kafkaObjs != nil && len(kafkaObjs) > 0 {
		kafkaDB.InitObj(kafkaObjs[0])
	}
}

// DeInit 断开数据库连接
func DeInit() {
	mongoDB.Disconnect()
	mysqlDB.Disconnect()
	redisDB.Disconnect()
	kafkaDB.Disconnect()
}
