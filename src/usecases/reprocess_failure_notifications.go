package usecases

import (
	"fmt"
	"github.com/rodrinoblega/notification_handler/src/entities"
	"log"
	"time"
)

type ReprocessFailureNotificationUseCase struct {
	Repo      NotificationRepository
	ProviderA NotificationProvider
	ProviderB NotificationProvider
}

func NewReprocessFailureNotificationUseCase(repo NotificationRepository, providerA NotificationProvider, providerB NotificationProvider) *ReprocessFailureNotificationUseCase {
	return &ReprocessFailureNotificationUseCase{Repo: repo, ProviderA: providerA, ProviderB: providerB}
}

func (sn *ReprocessFailureNotificationUseCase) Reprocess() {
	notifications, err := sn.Repo.GetFailedNotifications()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, notification := range notifications {
		log.Printf("Processing notification %s with content %s", notification.ID, notification.Content)

		notification.Status = "retrying_async"
		sn.Repo.UpdateNotification(&notification)

		log.Printf("Trying provider A")
		if sn.sendWithRetries(&notification, sn.ProviderA) {
			sn.Repo.UpdateNotification(&notification)
			continue
		}

		log.Printf("Trying provider B")
		if sn.sendWithRetries(&notification, sn.ProviderB) {
			sn.Repo.UpdateNotification(&notification)
			continue
		}

		log.Printf("Both providers failed")
		notification.Status = "permanent_failure_async"
		sn.Repo.UpdateNotification(&notification)
	}

}

func (sn *ReprocessFailureNotificationUseCase) sendWithRetries(notification *entities.Notification, provider NotificationProvider) bool {
	retries := 0

	for retries < MaxRetries {
		err := provider.Send(notification.Content)
		if err == nil {
			notification.Status = "sent_async"
			return true
		}

		log.Printf("Error sending notification %s: %v\n", notification.ID, err)

		notification.Status = "failed_async"
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
