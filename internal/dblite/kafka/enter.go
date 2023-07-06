package kafka

import (
	"context"
	"errors"
	"github.com/Shopify/sarama"
	"github.com/loebfly/ezgin/ezlogs"
	"strings"
)

var client = Client{}

type Client struct{}

func InitObj(obj EZGinKafka) {
	ezlogs.CDebug("KAFKA", "初始化")
	config.initObj(obj)
	err := ctl.initConnect()
	if err != nil {
		ezlogs.CError("KAFKA", "初始化失败: {}", err.Error())
	} else {
		ezlogs.CInfo("KAFKA", "初始化成功")
	}
	// ctl.addCheckTicker()
}

func GetDB() Client {
	return client
}

func Disconnect() {
	ctl.disconnect()
}

// GetClient 获取客户端
func (Client) GetClient() sarama.Client {
	return ctl.client
}

// GetAsyncProducer 获取异步生产者
func (Client) GetAsyncProducer() (sarama.AsyncProducer, error) {
	if err := ctl.tryConnect(); err != nil {
		return nil, err
	}
	return sarama.NewAsyncProducerFromClient(ctl.client)
}

// GetSyncProducer 获取同步生产者
func (c Client) GetSyncProducer() (sarama.SyncProducer, error) {
	if err := ctl.tryConnect(); err != nil {
		return nil, err
	}
	return sarama.NewSyncProducerFromClient(ctl.client)
}

// GetConsumer 获取消费者
func (c Client) GetConsumer() (sarama.Consumer, error) {
	if err := ctl.tryConnect(); err != nil {
		return nil, err
	}
	return sarama.NewConsumerFromClient(ctl.client)
}

// GetConsumerGroup 获取消费者组
func (c Client) GetConsumerGroup(id string) (sarama.ConsumerGroup, error) {
	if err := ctl.tryConnect(); err != nil {
		return nil, err
	}
	return sarama.NewConsumerGroupFromClient(id, ctl.client)
}

// GetClusterAdmin 获取集群管理
func (c Client) GetClusterAdmin() (sarama.ClusterAdmin, error) {
	if err := ctl.tryConnect(); err != nil {
		return nil, err
	}
	return sarama.NewClusterAdminFromClient(ctl.client)
}

// CreateTopic 创建主题
func (c Client) CreateTopic(topic string) error {
	admin, err := c.GetClusterAdmin()
	if err != nil {
		return err
	}

	err = admin.CreateTopic(topic, &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false)
	if err != nil && !strings.Contains(err.Error(), "Topic with this name already exists") {
		ezlogs.CError("KAFKA", "创建topic失败: {}", err.Error())
		return err
	}
	return nil
}

// InputMsgForTopic 向主题中输入消息
func (c Client) InputMsgForTopic(topic string, message ...string) error {
	if len(message) == 0 {
		return errors.New("message 为空，无法输入")
	}

	err := c.CreateTopic(topic)
	if err != nil {
		return err
	}

	producer, err := c.GetAsyncProducer()
	if err != nil {
		return err
	}

	for _, content := range message {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(content),
		}
		producer.Input() <- msg
		ezlogs.CDebug("KAFKA", "成功向{}主题发送了消息:{}", topic, content)
	}
	return nil
}

// ListenTopicForGroupId 按组ID监听主题
func (c Client) ListenTopicForGroupId(topic, groupId string, handler func(msg string) error) {
	go func() {
		for {
			consumerGroup, err := c.GetConsumerGroup(groupId)
			if err != nil {
				ezlogs.CError("KAFKA", "消费者组{}监听{}主题失败: {}", groupId, topic, err.Error())
				continue
			}

			go func() {
				for err2 := range consumerGroup.Errors() {
					ezlogs.CError("KAFKA", "消费者组{}监听{}主题出错: {}", groupId, topic, err2.Error())
				}
			}()
			err3 := consumerGroup.Consume(context.Background(), []string{topic}, &msgConsumerGroupHandler{
				handler: handler,
			})
			if err3 != nil {
				ezlogs.CError("KAFKA", "消费{}组的{}主题失败: {}", groupId, topic, err3.Error())
			}
		}
	}()
}
