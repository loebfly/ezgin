package mysql

import (
	"errors"
	"fmt"
	"github.com/loebfly/ezgin/ezlogs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var ctl = new(control)

type control struct {
	dbMap map[string]*gorm.DB
}

func (c *control) initConnect() error {
	if c.dbMap == nil {
		c.dbMap = make(map[string]*gorm.DB)
	}
	for _, v := range config.Objs {
		gormDB, err := gorm.Open(mysql.Open(v.Url), &gorm.Config{})
		if err != nil {
			return err
		}
		sqlDB, err := gormDB.DB()
		if err != nil {
			return err
		}

		sqlDB.SetMaxOpenConns(v.Pool.Max)
		sqlDB.SetMaxIdleConns(v.Pool.Idle)
		sqlDB.SetConnMaxIdleTime(time.Duration(v.Pool.Timeout.Idle) * time.Second)
		sqlDB.SetConnMaxLifetime(time.Duration(v.Pool.Timeout.Life) * time.Minute)
		if v.Debug {
			c.dbMap[v.Tag] = gormDB.Debug()
		} else {
			c.dbMap[v.Tag] = gormDB
		}
	}
	return nil
}

func (c *control) tryConnect(tag string) error {
	if db, ok := c.dbMap[tag]; ok {
		if db != nil {
			sqlDB, err := db.DB()
			if err == nil {
				err = sqlDB.Ping()
				if err == nil {
					return nil
				}
			}
		}
	}
	for _, v := range config.Objs {
		if v.Tag == tag {
			gormDB, err := gorm.Open(mysql.Open(v.Url), &gorm.Config{})
			if err != nil {
				return err
			}
			sqlDB, err := gormDB.DB()
			if err != nil {
				return err
			}

			sqlDB.SetMaxOpenConns(v.Pool.Max)
			sqlDB.SetMaxIdleConns(v.Pool.Idle)
			sqlDB.SetConnMaxIdleTime(time.Duration(v.Pool.Timeout.Idle) * time.Second)
			sqlDB.SetConnMaxLifetime(time.Duration(v.Pool.Timeout.Life) * time.Minute)
			if v.Debug {
				c.dbMap[v.Tag] = gormDB.Debug()
			} else {
				c.dbMap[v.Tag] = gormDB
			}
			return nil
		}
	}
	return errors.New(fmt.Sprintf("未找到%s对应的Mysql数据库", tag))
}

func (c *control) disconnect() {
	for _, v := range c.dbMap {
		db, _ := v.DB()
		_ = db.Close()
	}
	c.dbMap = nil
}

func (c *control) retry() {
	for k := range c.dbMap {
		err := c.tryConnect(k)
		if err != nil {
			ezlogs.CError("MYSQL", "{} 对应的Mysql数据库重连失败: {}", k, err.Error())
		}
	}
}

func (c *control) addCheckTicker() {
	//设置定时任务自动检查
	ticker := time.NewTicker(time.Minute * 30)
	go func(c *control) {
		for range ticker.C {
			c.retry()
		}
	}(c)
}

func (c *control) getDB(findName ...string) (*gorm.DB, error) {
	key := ""
	if len(findName) == 0 {
		if len(config.Objs) > 0 {
			key = config.Objs[0].Tag
		} else {
			return nil, errors.New("未配置Mysql数据库")
		}

	} else {
		key = findName[0]
	}
	//if db, ok := c.dbMap[key]; ok {
	//	return db, nil
	//}
	err := c.tryConnect(key)
	if err != nil {
		return nil, err
	}
	return c.dbMap[key], nil
}
