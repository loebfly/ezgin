package mongo

import (
	"errors"
	"fmt"
	"github.com/loebfly/ezgin/internal/logs"
	"gopkg.in/mgo.v2"
	"time"
)

var ctl = new(control)

type control struct {
	dbMap map[string]*mgo.Session
}

func (c *control) initConnect() error {
	if c.dbMap == nil {
		c.dbMap = make(map[string]*mgo.Session)
	}
	for _, v := range config.Objs {
		session, err := mgo.Dial(v.Url)
		if err != nil {
			return err
		}
		session.SetPoolLimit(v.PoolMax)
		session.SetMode(mgo.Monotonic, true)
		c.dbMap[v.Tag] = session
	}
	return nil
}

func (c *control) tryConnect(tag string) error {
	if db, ok := c.dbMap[tag]; ok {
		if db != nil {
			err := db.Ping()
			if err == nil {
				return nil
			}
		}
	}
	for _, v := range config.Objs {
		if v.Tag == tag {
			session, err := mgo.Dial(v.Url)
			if err != nil {
				return err
			}
			session.SetPoolLimit(v.PoolMax)
			session.SetMode(mgo.Monotonic, true)
			c.dbMap[v.Tag] = session
			return nil
		}
	}
	return errors.New(fmt.Sprintf("未找到%s对应的Mongo数据库", tag))
}

func (c *control) disconnect() {
	for _, db := range c.dbMap {
		if db != nil {
			db.Close()
			db = nil
		}
	}
}

func (c *control) retry() {
	for k := range c.dbMap {
		err := c.tryConnect(k)
		if err != nil {
			logs.Enter.CError("MONGO", "{} 对应的Mysql数据库重连失败, {}", k, err.Error())
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

// getDB 获取mongoDB
// 如果fineName为空且有已连接的数据库链接，则返回第一个链接，没有则返回错误
// 如果fineName不为空，则返回fineName对应的链接，如果没有则错误
func (c *control) getDB(tag ...string) (db *mgo.Database, returnDB func(db *mgo.Database), err error) {
	key := ""
	if len(tag) == 0 {
		if len(config.Objs) > 0 {
			key = config.Objs[0].Tag
		} else {
			return nil, c.returnDB, errors.New("未配置Mongo数据库")
		}
	} else {
		key = tag[0]
	}

	if v, ok := c.dbMap[key]; ok {
		return v.Copy().DB(key), c.returnDB, nil
	}

	err = c.tryConnect(key)
	if err != nil {
		return nil, c.returnDB, err
	}
	return c.dbMap[key].Copy().DB(key), c.returnDB, nil
}

func (c *control) returnDB(db *mgo.Database) {
	if db != nil {
		db.Session.Close()
	}
}
