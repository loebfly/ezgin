package dblite

import (
	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
	"gorm.io/gorm"
)

type enter int

const Enter = enter(0)

func Init(ymlPath string) error {
	return nil
}

func (enter) Mysql(name ...string) (db *gorm.DB, err error) {
	return nil, nil
}

func (enter) Mongo(name ...string) (db *mgo.Database, returnDB func(), err error) {
	return nil, nil, nil
}

func (enter) Redis(name ...string) (db *redis.Client, err error) {
	return nil, nil
}

func (enter) SafeExit() {

}
