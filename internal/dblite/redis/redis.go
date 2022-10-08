package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/loebfly/ezgin/internal/logs"
	"net"
	"time"
)

var ctl = new(control)

type control struct {
	dbMap map[string]*redis.Client
}

func (c *control) initConnect() error {
	if c.dbMap == nil {
		c.dbMap = make(map[string]*redis.Client)
	}
	for _, v := range config.Objs {
		addr := fmt.Sprintf("%s:%d", v.Host, v.Port)
		client := redis.NewClient(&redis.Options{
			Addr:         addr,
			Password:     v.Password,
			DB:           v.Database,
			PoolSize:     v.Pool.Max,
			MinIdleConns: v.Pool.Min,
			IdleTimeout:  time.Duration(v.Pool.Idle) * time.Minute,
			DialTimeout:  time.Duration(v.Pool.Timeout) * time.Second,
			Dialer: func() (net.Conn, error) {
				netDialer := &net.Dialer{
					Timeout:   5 * time.Second,
					KeepAlive: 5 * time.Minute,
				}
				return netDialer.Dial("tcp", addr)
			},
		})
		_, err := client.Ping().Result()
		if err != nil {
			return err
		}
		c.dbMap[v.Tag] = client
	}
	return nil
}

func (c *control) tryConnect(tag string) error {
	if db, ok := c.dbMap[tag]; !ok {
		if db != nil {
			_, err := db.Ping().Result()
			if err == nil {
				return nil
			}
		}
	}

	for _, v := range config.Objs {
		if v.Tag == tag {
			addr := fmt.Sprintf("%s:%d", v.Host, v.Port)
			client := redis.NewClient(&redis.Options{
				Addr:         addr,
				Password:     v.Password,
				DB:           v.Database,
				PoolSize:     v.Pool.Max,
				MinIdleConns: v.Pool.Min,
				IdleTimeout:  time.Duration(v.Pool.Idle) * time.Minute,
				DialTimeout:  time.Duration(v.Pool.Timeout) * time.Second,
				Dialer: func() (net.Conn, error) {
					netDialer := &net.Dialer{
						Timeout:   5 * time.Second,
						KeepAlive: 5 * time.Minute,
					}
					return netDialer.Dial("tcp", addr)
				},
			})
			_, err := client.Ping().Result()
			if err != nil {
				return err
			}
			c.dbMap[v.Tag] = client
			return nil
		}
	}
	return errors.New(fmt.Sprintf("未找到%s对应的Redis数据库", tag))
}

func (c *control) disconnect() {
	for _, v := range c.dbMap {
		_ = v.Close()
	}
	c.dbMap = nil
}

func (c *control) retry() {
	for k := range c.dbMap {
		err := c.tryConnect(k)
		if err != nil {
			logs.Enter.CError("REDIS", "%s对应的Redis数据库重连失败: %s", k, err.Error())
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

func (c *control) getDB(tag ...string) (*redis.Client, error) {
	key := ""
	if len(tag) == 0 {
		key = config.Objs[0].Tag
	} else {
		key = tag[0]
	}
	if db, ok := c.dbMap[key]; ok {
		return db, nil
	}
	err := c.tryConnect(key)
	if err != nil {
		return nil, err
	}
	return c.dbMap[key], nil
}