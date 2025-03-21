package usecases

import (
	"fmt"
	"github.com/rodrinoblega/notification_handler/src/entities"
	"log"
	"time"
)

type ReprocessFailureNotificationUseCase struct {
	NotificationRepository NotificationRepository
	UserActionRepository   UserActionRepository
	ProviderA              NotificationProvider
	ProviderB              NotificationProvider
}

func NewReprocessFailureNotificationUseCase(repo NotificationRepository, userActionRepository UserActionRepository, providerA NotificationProvider, providerB NotificationProvider) *ReprocessFailureNotificationUseCase {
	return &ReprocessFailureNotificationUseCase{NotificationRepository: repo, UserActionRepository: userActionRepository, ProviderA: providerA, ProviderB: providerB}
}

func (sn *ReprocessFailureNotificationUseCase) Reprocess() {
	notifications, err := sn.NotificationRepository.GetFailedNotifications()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, notification := range notifications {
		log.Printf("Getting template for %s.", notification.Type)
		contentWithPlaceholders := sn.NotificationRepository.GetTemplateOfNotification(notification.Type)

		ua, err := sn.UserActionRepository.GetByID(notification.UserActionID)
		if err != nil {
			fmt.Println(err)
			continue
		}

		finalContent := replaceTemplatePlaceholders(contentWithPlaceholders, &notification, ua)
		notification.Content = finalContent

		log.Printf("Processing notification %s with content %s", notification.ID, notification.Content)

		notification.Status = "processing_async"
		sn.NotificationRepository.UpdateNotification(&notification)

		log.Printf("Trying provider A")
		if sn.sendWithRetries(&notification, sn.ProviderA) {
			sn.NotificationRepository.UpdateNotification(&notification)
			continue
		}

		log.Printf("Trying provider B")
		if sn.sendWithRetries(&notification, sn.ProviderB) {
			sn.NotificationRepository.UpdateNotification(&notification)
			continue
		}

		log.Printf("Both providers failed")
		notification.Status = "permanent_failure_async"
		sn.NotificationRepository.UpdateNotification(&notification)
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

		notification.Status = "retrying_async"
		notification.Retries++
		_ = sn.NotificationRepository.UpdateNotification(notification)

		retries++
		backoffDuration := BaseBackoff * time.Duration(1<<retries)
		log.Printf("Retrying notification %s in %v\n", notification.ID, backoffDuration)
		time.Sleep(backoffDuration)
	}

	log.Printf("Provider reached max retries for notification %s.", notification.ID)
	return false
}
