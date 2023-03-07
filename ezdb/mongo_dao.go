package ezdb

import (
	"errors"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/ezlogs"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDao mongo数据库操作
type MongoDao[E MongoNameSetter] struct {
	DBTag func() string
}

// MongoNameSetter mongo数据库操作
type MongoNameSetter interface {
	MongoName() string
}

// Insert 插入数据
func (receiver *MongoDao[E]) Insert(entity E) error {
	var db *mgo.Database
	var returnDB func(db *mgo.Database)
	var err error
	if receiver.DBTag != nil {
		db, returnDB, err = Mongo(receiver.DBTag())
	} else {
		db, returnDB, err = Mongo()
	}
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
	var db *mgo.Database
	var returnDB func(db *mgo.Database)
	var err error
	if receiver.DBTag != nil {
		db, returnDB, err = Mongo(receiver.DBTag())
	} else {
		db, returnDB, err = Mongo()
	}
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
	var db *mgo.Database
	var returnDB func(db *mgo.Database)
	var err error
	if receiver.DBTag != nil {
		db, returnDB, err = Mongo(receiver.DBTag())
	} else {
		db, returnDB, err = Mongo()
	}
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

// Pager 分页查询
func (receiver *MongoDao[E]) Pager(db *mgo.Database, query bson.M, sort []string, page, pageSize int) ([]E, engine.Page, error) {
	var e E
	var total int
	var err error
	var result []E
	var returnDB func(db *mgo.Database)
	if db == nil {
		if receiver.DBTag != nil {
			db, returnDB, err = Mongo(receiver.DBTag())
		} else {
			db, returnDB, err = Mongo()
		}
		if err != nil {
			ezlogs.Error("数据库连接失败: {}", err.Error())
			return nil, engine.Page{}, errors.New("数据库连接失败")
		}
		defer returnDB(db)
	}
	total, err = db.C(e.MongoName()).Find(query).Count()
	if err != nil {
		ezlogs.Error("数据库查询失败: {}", err.Error())
		return nil, engine.Page{}, errors.New("数据库查询失败")
	}
	err = db.C(e.MongoName()).Find(query).Sort(sort...).Skip((page - 1) * pageSize).Limit(pageSize).All(&result)
	if err != nil {
		ezlogs.Error("数据库查询失败: {}", err.Error())
		return nil, engine.Page{}, errors.New("数据库查询失败")
	}

	var count int
	if total%pageSize == 0 {
		count = total / pageSize
	} else {
		count = total/pageSize + 1
	}

	return result, engine.Page{
		Total: total,
		Index: page,
		Size:  pageSize,
		Count: count,
	}, nil
}
