package kafkaclient

import "github.com/IBM/sarama"

var (
	TOPICS    = []string{"articles.new"}
	Version   = sarama.V3_9_0_0
	KAFKA_URL = "localhost:9092"
	BROKERS   = []string{KAFKA_URL}
)

func KafkaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V3_9_0_0
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	return config
}
