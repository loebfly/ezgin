package ezdb

import (
	"github.com/go-redis/redis"
	kafkaDB "github.com/loebfly/ezgin/internal/dblite/kafka"
	mongoDB "github.com/loebfly/ezgin/internal/dblite/mongo"
	mysqlDB "github.com/loebfly/ezgin/internal/dblite/mysql"
	redisDB "github.com/loebfly/ezgin/internal/dblite/redis"
	"gopkg.in/mgo.v2"
	"gorm.io/gorm"
)

// Mysql 获取mysql数据库
func Mysql(tag ...string) (db *gorm.DB, err error) {
	return mysqlDB.GetDB(tag...)
}

// Mongo 获取mongo数据库
func Mongo(tag ...string) (db *mgo.Database, returnDB func(db *mgo.Database), err error) {
	return mongoDB.GetDB(tag...)
}

// Redis 获取redis数据库
func Redis(tag ...string) (db *redis.Client, err error) {
	return redisDB.GetDB(tag...)
}

// Kafka 获取kafka数据库
func Kafka() kafkaDB.Client {
	return kafkaDB.GetDB()
}
