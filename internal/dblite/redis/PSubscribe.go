package redis

import (
	"github.com/go-redis/redis"
	"github.com/loebfly/ezgin/ezlogs"
	"time"
)

// PSubOperator 订阅操作
type PSubOperator struct {
	dbTag    []string
	channels []string
}

// SetDBTag 设置数据库标签
func (receiver *PSubOperator) SetDBTag(dbTag ...string) *PSubOperator {
	receiver.dbTag = dbTag
	return receiver
}

// SetChannels 设置订阅的频道
func (receiver *PSubOperator) SetChannels(channels ...string) *PSubOperator {
	receiver.channels = channels
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
