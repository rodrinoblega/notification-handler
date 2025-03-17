package usecases

import (
	"github.com/rodrinoblega/notification_handler/src/entities"
	"log"
	"time"
)

const (
	MaxRetries  = 3
	BaseBackoff = 2 * time.Second // Tiempo base para el backoff exponencial
)

type NotificationRepository interface {
	Save(n *entities.Notification) error
	UpdateStatus(id string, status string) (*entities.Notification, error)
	UpdateRetries(id string, retries int) (*entities.Notification, error)
	UpdateNotification(n *entities.Notification) error
	GetFailedNotifications() ([]entities.Notification, error)
}

type NotificationProvider interface {
	Send(content string) error
}

type SaveNotificationUseCase struct {
	Repo      NotificationRepository
	ProviderA NotificationProvider
	ProviderB NotificationProvider
}

func NewSaveNotificationUseCase(repo NotificationRepository, providerA NotificationProvider, providerB NotificationProvider) *SaveNotificationUseCase {
	return &SaveNotificationUseCase{Repo: repo, ProviderA: providerA, ProviderB: providerB}
}

func (sn *SaveNotificationUseCase) CreateNotification(notification entities.Notification) error {
	notification.ID = time.Now().Format("20060102150405")
	notification.CreatedAt = time.Now()
	notification.UpdatedAt = time.Now()
	notification.Status = "processing"

	if err := sn.Repo.Save(&notification); err != nil {
		return err
	}

	log.Printf("Trying provider A")
	if sn.sendWithRetries(&notification, sn.ProviderA) {
		return sn.Repo.UpdateNotification(&notification)
	}

	log.Printf("Trying provider B")
	if sn.sendWithRetries(&notification, sn.ProviderB) {
		return sn.Repo.UpdateNotification(&notification)
	}

	log.Printf("Both providers failed")
	notification.Status = "permanent_failure"
	return sn.Repo.UpdateNotification(&notification)
}

func (sn *SaveNotificationUseCase) sendWithRetries(notification *entities.Notification, provider NotificationProvider) bool {
	retries := 0

	for retries < MaxRetries {
		err := provider.Send(notification.Content)
		if err == nil {
			notification.Status = "sent"
			return true
		}

		log.Printf("Error sending notification %s: %v\n", notification.ID, err)

		notification.Status = "failed"
		notification.Retries++
		_ = sn.Repo.UpdateNotification(notification)

		retries++
		backoffDuration := BaseBackoff * time.Duration(1<<retries)
		log.Printf("Retrying notification %s in %v\n", notification.ID, backoffDuration)
		time.Sleep(backoffDuration)
	}

	log.Printf("Provider reached max retries for notification %s.", notification.ID)
	return false
}
