package ezdb

import (
	"errors"
	"github.com/loebfly/ezgin/ezlogs"
	"gopkg.in/mgo.v2/bson"
)

// MongoDao mongo数据库操作
type MongoDao[E MongoNameSetter] struct{}

// MongoNameSetter mongo数据库操作
type MongoNameSetter interface {
	MongoName() string
}

// Insert 插入数据
func (receiver *MongoDao[E]) Insert(entity E) error {
	db, returnDB, err := Mongo()
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return errors.New("数据库连接失败")
	}
	defer returnDB(db)
	err = db.C(entity.MongoName()).Insert(entity)
	if err != nil {
		ezlogs.Error("数据库插入失败: {}", err.Error())
		return errors.New("数据库插入失败")
	}
	return nil
}

// RemoveId 删除数据
func (receiver *MongoDao[E]) RemoveId(id bson.ObjectId) error {
	db, returnDB, err := Mongo()
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return errors.New("数据库连接失败")
	}
	defer returnDB(db)
	var e E
	err = db.C(e.MongoName()).RemoveId(id)
	if err != nil {
		ezlogs.Error("数据库删除失败: {}", err.Error())
		return errors.New("数据库删除失败")
	}
	return nil
}

// UpdateId 更新数据
func (receiver *MongoDao[E]) UpdateId(id bson.ObjectId, entity E) error {
	db, returnDB, err := Mongo()
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return errors.New("数据库连接失败")
	}
	defer returnDB(db)
	err = db.C(entity.MongoName()).UpdateId(id, entity)
	if err != nil {
		ezlogs.Error("数据库更新失败: {}", err.Error())
		return errors.New("数据库更新失败")
	}
	return nil
}
