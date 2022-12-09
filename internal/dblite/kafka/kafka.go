package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/loebfly/ezgin/ezlogs"
	"strings"
	"time"
)

var ctl = new(control)

type control struct {
	client sarama.Client
}

func (c *control) initConnect() error {
	saramaCfg := sarama.NewConfig()
	switch config.Obj.Ack {
	case "no":
		saramaCfg.Producer.RequiredAcks = sarama.NoResponse
	case "local":
		saramaCfg.Producer.RequiredAcks = sarama.WaitForLocal
	case "all":
		saramaCfg.Producer.RequiredAcks = sarama.WaitForAll
	default:
		saramaCfg.Producer.RequiredAcks = sarama.WaitForAll
	}

	saramaCfg.Consumer.Offsets.AutoCommit.Enable = config.Obj.AutoCommit

	switch config.Obj.Partitioner {
	case "hash":
		saramaCfg.Producer.Partitioner = sarama.NewHashPartitioner
	case "random":
		saramaCfg.Producer.Partitioner = sarama.NewRandomPartitioner
	case "round-robin":
		saramaCfg.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	default:
		saramaCfg.Producer.Partitioner = sarama.NewHashPartitioner
	}

	version, _ := sarama.ParseKafkaVersion(config.Obj.Version)
	saramaCfg.Version = version

	servers := strings.Split(config.Obj.Servers, ",")
	client, err := sarama.NewClient(servers, saramaCfg)
	if err != nil {
		return errors.New("kafka.NewClient: " + err.Error())
	}
	if len(client.Brokers()) == 0 {
		return errors.New("kafka连接失败, 请检查配置, 无法连接到kafka服务器")
	} else if strings.Contains(client.Brokers()[0].Addr(), "127.0.0.1") {
		return errors.New("kafka连接失败, 请检查配置, 不能连接到127.0.0.1")
	}
	c.client = client
	return nil
}

func (c *control) tryConnect() error {
	if c.client == nil || c.client.Closed() {
		return c.initConnect()
	}
	return nil
}

func (c *control) disconnect() {
	if c.client != nil {
		err := c.client.Close()
		if err != nil {
			ezlogs.CError("KAFKA", "断开连接错误: {}", err.Error())
		}
	}
}

func (c *control) retry() {
	err := c.tryConnect()
	if err != nil {
		ezlogs.CError("KAFKA", "重试连接失败: {}", err.Error())
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

func (c *control) getDB() *control {
	return c
}

type msgConsumerGroupHandler struct {
	handler func(string) error
}

func (msgConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (msgConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h msgConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		//fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		err := h.handler(string(msg.Value))
		if err != nil {
			ezlogs.CError("KAFKA", "消费消息失败: {}", err.Error())
		}
		sess.MarkMessage(msg, "")
	}
	return nil
}
