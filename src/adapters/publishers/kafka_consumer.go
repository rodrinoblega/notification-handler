package publishers

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rodrinoblega/notification_handler/src/entities"
	"github.com/rodrinoblega/notification_handler/src/usecases"
	"log"
)

type KafkaConsumer struct {
	Consumer         *kafka.Consumer
	Topic            string
	SaveNotification *usecases.SaveNotificationUseCase
}

func (kc *KafkaConsumer) Subscribe() {
	err := kc.Consumer.Subscribe(kc.Topic, nil)
	if err != nil {
		log.Fatalf("Error while subscribig to topic: %v", err)
	}

	fmt.Println("Kafka Consumer initialized. Waiting for messages...")

	for {
		msg, err := kc.Consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Received message: %s\n", string(msg.Value))

			var userAction entities.UserAction
			if err := json.Unmarshal(msg.Value, &userAction); err != nil {
				log.Println("Error parsing message:", err)
				continue
			}

			if err := kc.SaveNotification.CreateNotification(&userAction); err != nil {
				log.Println("Error saving notification:", err)
			} else {
				log.Println("Notification saved successfully!")
			}
		} else {
			log.Println("Consumer error:", err)
		}
	}

}
