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

// Mysql 获取mysql数据库, tag为空时返回第一个数据库, tag 多个只取第一个
func Mysql(tag ...string) (db *gorm.DB, err error) {
	return mysqlDB.GetDB(tag...)
}

// GetMysqlAllTags 获取所有mysql数据库标签
func GetMysqlAllTags() []string {
	return mysqlDB.GetAllTags()
}

// Mongo 获取mongo数据库, tag为空时返回第一个数据库, tag 多个只取第一个
func Mongo(tag ...string) (db *mgo.Database, returnDB func(db *mgo.Database), err error) {
	return mongoDB.GetDB(tag...)
}

// GetMongoAllTags 获取所有mongo数据库标签
func GetMongoAllTags() []string {
	return mongoDB.GetAllTags()
}

// Redis 获取redis数据库, tag为空时返回第一个数据库, tag 多个只取第一个
func Redis(tag ...string) (db redis.UniversalClient, err error) {
	return redisDB.GetDB(tag...)
}

// GetRedisAllTags 获取所有redis数据库标签
func GetRedisAllTags() []string {
	return redisDB.GetAllTags()
}

// Kafka 获取kafka数据库
func Kafka() kafkaDB.Client {
	return kafkaDB.GetDB()
}
