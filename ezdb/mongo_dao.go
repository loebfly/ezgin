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

// GetDB 获取Mongo数据库
func (receiver *MongoDao[E]) GetDB() (*mgo.Database, func(db *mgo.Database), error) {
	var db *mgo.Database
	var returnDB func(db *mgo.Database)
	var err error
	if receiver.DBTag != nil {
		db, returnDB, err = Mongo(receiver.DBTag())
	} else {
		db, returnDB, err = Mongo()
	}
	if err != nil {
		ezlogs.Error("Mongo数据库连接失败: {}", err.Error())
		return nil, nil, errors.New("数据库连接失败")
	}
	return db, returnDB, nil
}

// Insert 插入数据
func (receiver *MongoDao[E]) Insert(entity E) error {
	db, returnDB, err := receiver.GetDB()
	if err != nil {
		return err
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
	db, returnDB, err := receiver.GetDB()
	if err != nil {
		return err
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
	db, returnDB, err := receiver.GetDB()
	if err != nil {
		return err
	}
	defer returnDB(db)
	err = db.C(entity.MongoName()).UpdateId(id, entity)
	if err != nil {
		ezlogs.Error("数据库更新失败: {}", err.Error())
		return errors.New("数据库更新失败")
	}
	return nil
}

// All 查询所有数据
func (receiver *MongoDao[E]) All(query bson.M, sort ...string) ([]E, error) {
	db, returnDB, err := receiver.GetDB()
	if err != nil {
		return nil, err
	}
	defer returnDB(db)
	var e E
	find := db.C(e.MongoName()).Find(query)
	if sort != nil {
		find = find.Sort(sort...)
	}

	var result = make([]E, 0)
	err = find.All(&result)
	if err != nil {
		ezlogs.Error("数据库查询失败: {}", err.Error())
		return nil, errors.New("数据库查询失败")
	}
	return result, nil
}

// One 查询单条数据
func (receiver *MongoDao[E]) One(query bson.M) (*E, error) {
	db, returnDB, err := receiver.GetDB()
	if err != nil {
		return nil, err
	}
	defer returnDB(db)

	var e E
	err = db.C(e.MongoName()).Find(query).One(&e)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		ezlogs.Error("数据库查询失败: {}", err.Error())
		return nil, errors.New("数据库查询失败")
	}

	return &e, nil
}

// Pager 分页查询 db
func (receiver *MongoDao[E]) Pager(query bson.M, page, pageSize int, sort ...string) ([]E, engine.Page, error) {
	db, returnDB, err := receiver.GetDB()
	if err != nil {
		return nil, engine.Page{}, err
	}
	defer returnDB(db)
	var e E
	find := db.C(e.MongoName()).Find(query)
	var total int
	total, err = find.Count()
	if err != nil {
		ezlogs.Error("数据库查询失败: {}", err.Error())
		return nil, engine.Page{}, errors.New("数据库查询失败")
	}

	if sort != nil {
		find = find.Sort(sort...)
	}

	var result = make([]E, 0)
	err = find.Skip((page - 1) * pageSize).Limit(pageSize).All(&result)
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
