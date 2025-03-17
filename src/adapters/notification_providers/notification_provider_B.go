package notification_providers

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

type NotificationProviderB struct{}

func NewNotificationProviderB() *NotificationProviderB {
	return &NotificationProviderB{}
}

func (np *NotificationProviderB) Send(content string) error {
	log.Printf("Simulating sending notification: %s\n", content)

	if rand.Float32() < 0.5 {
		return errors.New("simulated failure")
	}

	time.Sleep(2 * time.Second)
	log.Println("Notification sent successfully!")
	return nil
}
