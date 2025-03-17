package notification_providers

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

type NotificationProviderA struct{}

func NewNotificationProviderA() *NotificationProviderA {
	return &NotificationProviderA{}
}

func (np *NotificationProviderA) Send(content string) error {
	log.Printf("Simulating sending notification: %s\n", content)

	if rand.Float32() < 0.8 {
		return errors.New("simulated failure")
	}

	time.Sleep(2 * time.Second)
	log.Println("Notification sent successfully!")
	return nil
}
