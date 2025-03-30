package kafkaclient

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func KafkaSyncProducer() sarama.SyncProducer {
	config := KafkaConfig()
	producer, err := sarama.NewSyncProducer(BROKERS, config)
	if err != nil {
		panic(err)
	}
	return producer
}

func KafkaAsyncProducer() sarama.AsyncProducer {
	config := KafkaConfig()
	producer, err := sarama.NewAsyncProducer(BROKERS, config)
	if err != nil {
		panic(err)
	}
	return producer
}

func KafkaMessage(topic string, value string) *sarama.ProducerMessage {
	message := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(value)}
	return message
}

func ProduceAsyncMessage(producer sarama.AsyncProducer, message *sarama.ProducerMessage) error {
	producer.Input() <- message
	select {
	case success := <-producer.Successes():
		fmt.Printf("Message sent to partition %d at offset %d\n", success.Partition, success.Offset)
	case err := <-producer.Errors():
		log.Printf("Failed to send message: %v", err)
	}

	return nil
}
