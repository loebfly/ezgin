package kafka

import (
	"github.com/Shopify/sarama"
	"strings"
)

var ctl = new(control)

type control struct {
	client sarama.Client
}

func (ctl *control) initConnect() error {
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
		return err
	}
	ctl.client = client
	return nil
}
