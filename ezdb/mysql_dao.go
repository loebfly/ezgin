package ezdb

import (
	"errors"
	"github.com/loebfly/ezgin/engine"
	"github.com/loebfly/ezgin/ezlogs"
	"gorm.io/gorm"
)

type MysqlDao[E any] struct {
	DBTag func() string // 数据库标签
	debug bool          // 是否开启debug模式
}

// Debug 开启debug模式
func (receiver *MysqlDao[E]) Debug() *MysqlDao[E] {
	newDao := &MysqlDao[E]{
		DBTag: receiver.DBTag,
		debug: true,
	}
	return newDao
}

// Create 插入数据
func (receiver *MysqlDao[E]) Create(entity *E) error {
	var db *gorm.DB
	var err error
	if receiver.DBTag == nil {
		db, err = Mysql()
	} else {
		db, err = Mysql(receiver.DBTag())
	}
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return errors.New("数据库连接失败")
	}
	if receiver.debug {
		db = db.Debug()
	}
	err = db.Create(entity).Error
	if err != nil {
		ezlogs.Error("数据库插入失败: {}", err.Error())
		return errors.New("数据库插入失败")
	}
	return nil
}

// MultiCreate 插入多条数据
func (receiver *MysqlDao[E]) MultiCreate(entities []*E) error {
	var db *gorm.DB
	var err error
	if receiver.DBTag == nil {
		db, err = Mysql()
	} else {
		db, err = Mysql(receiver.DBTag())
	}
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return errors.New("数据库连接失败")
	}
	if receiver.debug {
		db = db.Debug()
	}
	err = db.Create(entities).Error
	if err != nil {
		ezlogs.Error("数据库插入失败: {}", err.Error())
		return errors.New("数据库插入失败")
	}
	return nil
}

// Delete 删除数据
func (receiver *MysqlDao[E]) Delete(entity E) error {
	var db *gorm.DB
	var err error
	if receiver.DBTag == nil {
		db, err = Mysql()
	} else {
		db, err = Mysql(receiver.DBTag())
	}
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return errors.New("数据库连接失败")
	}
	if receiver.debug {
		db = db.Debug()
	}
	var e E
	err = db.Where(entity).Delete(&e).Error
	if err != nil {
		ezlogs.Error("数据库删除失败: {}", err.Error())
		return errors.New("数据库删除失败")
	}
	return nil
}

// Updates 更新数据
func (receiver *MysqlDao[E]) Updates(entity *E) error {
	var db *gorm.DB
	var err error
	if receiver.DBTag == nil {
		db, err = Mysql()
	} else {
		db, err = Mysql(receiver.DBTag())
	}
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return errors.New("数据库连接失败")
	}
	if receiver.debug {
		db = db.Debug()
	}
	err = db.Updates(entity).Error
	if err != nil {
		ezlogs.Error("数据库更新失败: {}", err.Error())
		return errors.New("数据库更新失败")
	}
	return nil
}

// Save 保存数据
func (receiver *MysqlDao[E]) Save(entity *E) error {
	var db *gorm.DB
	var err error
	if receiver.DBTag == nil {
		db, err = Mysql()
	} else {
		db, err = Mysql(receiver.DBTag())
	}
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return errors.New("数据库连接失败")
	}
	if receiver.debug {
		db = db.Debug()
	}
	err = db.Save(entity).Error
	if err != nil {
		ezlogs.Error("数据库保存失败: {}", err.Error())
		return errors.New("数据库保存失败")
	}
	return nil
}

// All mysql动态查询数据
func (receiver *MysqlDao[E]) All(entity E) ([]E, error) {
	var db *gorm.DB
	var err error
	if receiver.DBTag == nil {
		db, err = Mysql()
	} else {
		db, err = Mysql(receiver.DBTag())
	}
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return nil, errors.New("数据库连接失败")
	}
	if receiver.debug {
		db = db.Debug()
	}
	var result = make([]E, 0)
	err = db.Where(entity).Find(&result).Error
	if err != nil {
		ezlogs.Error("数据库查询失败: {}", err.Error())
		return nil, errors.New("数据库查询失败")
	}
	return result, nil
}

// One mysql动态查询一条数据
func (receiver *MysqlDao[E]) One(entity E) (*E, error) {
	var db *gorm.DB
	var err error
	if receiver.DBTag == nil {
		db, err = Mysql()
	} else {
		db, err = Mysql(receiver.DBTag())
	}
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return nil, errors.New("数据库连接失败")
	}
	if receiver.debug {
		db = db.Debug()
	}
	var result *E
	err = db.Where(entity).Find(&result).Error
	if err != nil {
		ezlogs.Error("数据库查询失败: {}", err.Error())
		return nil, errors.New("数据库查询失败")
	}
	return result, nil
}

// Pager 分页查询, db为已经设置好的查询条件的db，page为当前页，pageSize为每页条数，返回值为查询结果，分页信息，错误信息
func (receiver *MysqlDao[E]) Pager(db *gorm.DB, page, pageSize int) ([]E, engine.Page, error) {
	var err error
	var result = make([]E, 0)
	var total int64
	offset := (page - 1) * pageSize
	err = db.Offset(offset).Limit(pageSize).Find(&result).Error
	if err != nil {
		ezlogs.Error("数据库查询失败: {}", err.Error())
		return nil, engine.Page{}, errors.New("数据库查询失败")
	}
	err = db.Count(&total).Error
	if err != nil {
		ezlogs.Error("数据库查询失败: {}", err.Error())
		return nil, engine.Page{}, errors.New("数据库查询失败")
	}

	var count int
	if total%int64(pageSize) == 0 {
		count = int(total) / pageSize
	} else {
		count = int(total)/pageSize + 1
	}

	return result, engine.Page{
		Total: int(total),
		Size:  pageSize,
		Index: page,
		Count: count,
	}, nil
}
