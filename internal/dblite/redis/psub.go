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

// Receive 接收
func (receiver *PSubOperator) Receive(handler func(dbTag *string, msg *redis.Message)) {
	if handler == nil {
		ezlogs.CError("REDIS", "PSubOperator err: handler is nil")
		return
	}
	rds, err := ctl.getDB(receiver.dbTag...)
	if err != nil {
		ezlogs.CError("REDIS", "PSubOperator err: {}", err.Error())
		return
	}
	pubSub := rds.PSubscribe(receiver.channels...)
	msgCn := pubSub.Channel()
	for msg := range msgCn {
		ezlogs.CDebug("REDIS", "PSubOperator receive Payload:{}", msg.Payload)
		var dbTag *string
		if len(receiver.dbTag) > 0 {
			dbTag = &receiver.dbTag[0]
		} else {
			dbTag = &config.Objs[0].Tag
		}
		handler(dbTag, msg)
	}
	_ = pubSub.Close()
}

// HoldReceive 保持接收
func (receiver *PSubOperator) HoldReceive(handler func(dbTag *string, msg *redis.Message)) {
	if handler == nil {
		ezlogs.CError("REDIS", "PSubOperator err: handler is nil")
		return
	}
	go func() {
		for {
			receiver.Receive(handler)
			time.Sleep(time.Second)
		}
	}()
}
