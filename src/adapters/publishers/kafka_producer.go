package publishers

import (
	"encoding/json"
	"github.com/rodrinoblega/notification_handler/src/entities"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	Producer *kafka.Producer
	Topic    string
}

func (kp *KafkaProducer) Publish(notification entities.Notification) error {
	msg, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	deliveryChan := make(chan kafka.Event)
	err = kp.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.Topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, deliveryChan)

	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		log.Printf("Message delivered to %v\n", m.TopicPartition)
	}

	return nil
}
