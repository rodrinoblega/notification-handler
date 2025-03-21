package messaging

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rodrinoblega/notification_handler/config"
	"github.com/rodrinoblega/notification_handler/src/adapters/publishers"
	"github.com/rodrinoblega/notification_handler/src/usecases"
)

func NewKafkaProducer(env *config.Config) (*publishers.KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": env.KafkaBrokers,
	})
	if err != nil {
		return nil, err
	}

	return &publishers.KafkaProducer{
		Producer: p,
		Topic:    env.KafkaTopic,
	}, nil
}

func NewKafkaConsumer(env *config.Config, useCase *usecases.ProcessNotificationUseCase) (*publishers.KafkaConsumer, error) {
	groupID := "notification-consumer-group"

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": env.KafkaBrokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	return &publishers.KafkaConsumer{
		Consumer:                   consumer,
		Topic:                      env.KafkaTopic,
		ProcessNotificationUseCase: useCase,
	}, nil

}
