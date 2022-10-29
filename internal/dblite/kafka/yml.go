package kafka

type Yml struct {
	EZGinKafka EZGinKafka `koanf:"ezgin_kafka"`
}

type EZGinKafka struct {
	Servers     string `koanf:"servers"`     // ip:port 集群多个服务器之间用逗号分隔
	Ack         string `koanf:"ack"`         // ack模式 no,local,all, 默认all
	AutoCommit  bool   `koanf:"auto_commit"` // 是否自动提交, 默认true
	Partitioner string `koanf:"partitioner"` // 分区选择模式 hash,random,round-robin, 默认hash
	Version     string `koanf:"version"`     // kafka版本 默认2.8.1
}
