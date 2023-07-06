package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/loebfly/ezgin/ezlogs"
	"strconv"
	"time"
)

// PSubOperator 订阅操作
type PSubOperator struct {
	dbTag    []string
	channels []string
}

// SetChannels 设置订阅的频道
func (receiver *PSubOperator) SetChannels(channels ...string) *PSubOperator {
	receiver.channels = channels
	return receiver
}

// AddKeyExpiredChannel 添加键过期的频道
func (receiver *PSubOperator) AddKeyExpiredChannel() *PSubOperator {
	dbNum := "*" // 默认为任意数据库
	if len(config.Objs) > 0 {
		if len(receiver.dbTag) > 0 {
			for _, obj := range config.Objs {
				if obj.Tag == receiver.dbTag[0] {
					dbNum = strconv.Itoa(obj.GetDB())
				}
			}
		} else {
			dbNum = strconv.Itoa(config.Objs[0].GetDB())
		}
	}

	receiver.channels = append(receiver.channels, fmt.Sprintf("__keyevent@%s__:expired", dbNum))
	return receiver
}

// Run 运行
func (receiver *PSubOperator) Run(handler func(msg *redis.Message)) {
	if handler == nil {
		ezlogs.Error("REDIS", "PSubOperator.Run: handler is nil")
		return
	}
	rds, err := ctl.getDB(receiver.dbTag...)
	if err != nil {
		ezlogs.Error("REDIS", "PSubOperator.Do: {}", err.Error())
		return
	}
	pubSub := rds.Subscribe(receiver.channels...)
	msgCn := pubSub.Channel()
	for msg := range msgCn {
		ezlogs.Debug("REDIS", "收到Redis消息:{}", msg.Payload)
		handler(msg)
	}
	_ = pubSub.Close()
}

// HoldRun 持续运行
func (receiver *PSubOperator) HoldRun(handler func(msg *redis.Message)) {
	if handler == nil {
		ezlogs.Error("REDIS", "PSubOperator.Run: handler is nil")
		return
	}
	go func() {
		for {
			receiver.Run(handler)
			time.Sleep(time.Second)
		}
	}()
}
