package kafka

import (
	"context"
	"errors"
	"github.com/Shopify/sarama"
	"github.com/loebfly/ezgin/internal/logs"
)

var client = Client{}

type Client struct {
	consumerGroups map[string]sarama.ConsumerGroup
}

func InitObj(obj EZGinKafka) {
	logs.Enter.CDebug("KAFKA", "初始化")
	config.initObj(obj)
	err := ctl.initConnect()
	if err != nil {
		logs.Enter.CError("KAFKA", "初始化失败: {}", err.Error())
	}
	logs.Enter.CInfo("KAFKA", "初始化成功")
	ctl.addCheckTicker()
}

func GetDB() Client {
	return client
}

func Disconnect() {
	if len(client.consumerGroups) > 0 {
		for groupId, group := range client.consumerGroups {
			err := group.Close()
			if err != nil {
				logs.Enter.CError("kAFKA", "关闭{}消费者组失败: {}", groupId, err.Error())
				return
			}
		}
	}
	ctl.disconnect()
}

// GetClient 获取客户端
func (Client) GetClient() sarama.Client {
	return ctl.client
}

// GetAsyncProducer 获取异步生产者
func (Client) GetAsyncProducer() (sarama.AsyncProducer, error) {
	if ctl.client == nil {
		return nil, errors.New("kafka client is nil")
	}
	return sarama.NewAsyncProducerFromClient(ctl.client)
}

// GetSyncProducer 获取同步生产者
func (c Client) GetSyncProducer() (sarama.SyncProducer, error) {
	if ctl.client == nil {
		return nil, errors.New("kafka client is nil")
	}
	return sarama.NewSyncProducerFromClient(ctl.client)
}

// GetConsumer 获取消费者
func (c Client) GetConsumer() (sarama.Consumer, error) {
	if ctl.client == nil {
		return nil, errors.New("kafka client is nil")
	}
	return sarama.NewConsumerFromClient(ctl.client)
}

// GetConsumerGroup 获取消费者组
func (c Client) GetConsumerGroup(id string) (sarama.ConsumerGroup, error) {
	if ctl.client == nil {
		return nil, errors.New("kafka client is nil")
	}
	return sarama.NewConsumerGroupFromClient(id, ctl.client)
}

// GetClusterAdmin 获取集群管理
func (c Client) GetClusterAdmin() (sarama.ClusterAdmin, error) {
	if ctl.client == nil {
		return nil, errors.New("kafka client is nil")
	}
	return sarama.NewClusterAdminFromClient(ctl.client)
}

// CreateTopic 创建主题
func (c Client) CreateTopic(topic string) error {
	admin, err := c.GetClusterAdmin()
	if err != nil {
		return err
	}
	defer func(admin sarama.ClusterAdmin) {
		closeErr := admin.Close()
		if closeErr != nil {
			logs.Enter.CError("kAFKA", "关闭ClusterAdmin失败: {}", closeErr.Error())
		}
	}(admin)

	err = admin.CreateTopic(topic, &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false)
	if err != nil {
		logs.Enter.CError("kAFKA", "创建topic失败: {}", err.Error())
		return err
	}
	return nil
}

// InputMsgForTopic 向主题中输入消息
func (c Client) InputMsgForTopic(topic string, message ...string) error {
	if len(message) == 0 {
		return errors.New("message 为空，无法输入")
	}
	producer, err := c.GetSyncProducer()
	if err != nil {
		return err
	}
	defer func(producer sarama.SyncProducer) {
		closeErr := producer.Close()
		if closeErr != nil {
			logs.Enter.CError("kAFKA", "关闭SyncProducer失败: {}", closeErr.Error())
		}
	}(producer)

	for _, msg := range message {
		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(msg),
		})
		if err != nil {
			return errors.New("输入{" + msg + "}消息失败")
		}
	}
	return nil
}

// ListenTopicForGroupId 按组ID监听主题
func (c Client) ListenTopicForGroupId(topic, groupId string, handler func(msg string) error) error {
	consumerGroup, err := c.GetConsumerGroup(groupId)
	if err != nil {
		return err
	}

	if c.consumerGroups == nil {
		c.consumerGroups = make(map[string]sarama.ConsumerGroup)
	}
	c.consumerGroups[groupId] = consumerGroup

	err = consumerGroup.Consume(context.Background(), []string{topic}, &msgConsumerGroupHandler{
		handler: handler,
	})
	if err != nil {
		logs.Enter.CError("kAFKA", "监听{}组的{}主题失败: {}", groupId, topic, err.Error())
		return err
	}

	return nil
}
