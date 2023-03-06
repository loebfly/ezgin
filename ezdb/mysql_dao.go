package ezdb

import (
	"errors"
	"github.com/loebfly/ezgin/ezlogs"
)

type MysqlDao[E any] struct{}

// Create 插入数据
func (receiver *MysqlDao[E]) Create(entity *E) error {
	db, err := Mysql()
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return errors.New("数据库连接失败")
	}
	err = db.Create(entity).Error
	if err != nil {
		ezlogs.Error("数据库插入失败: {}", err.Error())
		return errors.New("数据库插入失败")
	}
	return nil
}

// Delete 删除数据
func (receiver *MysqlDao[E]) Delete(entity E) error {
	db, err := Mysql()
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return errors.New("数据库连接失败")
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
	db, err := Mysql()
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return errors.New("数据库连接失败")
	}
	err = db.Debug().Updates(entity).Error
	if err != nil {
		ezlogs.Error("数据库更新失败: {}", err.Error())
		return errors.New("数据库更新失败")
	}
	return nil
}

// Save 保存数据
func (receiver *MysqlDao[E]) Save(entity *E) error {
	db, err := Mysql()
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return errors.New("数据库连接失败")
	}
	err = db.Debug().Save(entity).Error
	if err != nil {
		ezlogs.Error("数据库保存失败: {}", err.Error())
		return errors.New("数据库保存失败")
	}
	return nil
}

// All mysql动态查询数据
func (receiver *MysqlDao[E]) All(entity E) ([]E, error) {
	db, err := Mysql()
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return nil, errors.New("数据库连接失败")
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
	db, err := Mysql()
	if err != nil {
		ezlogs.Error("数据库连接失败: {}", err.Error())
		return nil, errors.New("数据库连接失败")
	}
	var result *E
	err = db.Where(entity).Find(&result).Error
	if err != nil {
		ezlogs.Error("数据库查询失败: {}", err.Error())
		return nil, errors.New("数据库查询失败")
	}
	return result, nil
}
