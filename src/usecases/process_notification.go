package usecases

import (
	"fmt"
	"github.com/rodrinoblega/notification_handler/src/entities"
	"log"
	"strings"
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
	GetNotifications(userID string, notificationType string, startDate, endDate time.Time, limit, offset int) ([]entities.Notification, error)
	GetTemplateOfNotification(notificationType string) string
}

type NotificationProvider interface {
	Send(content string) error
}

type ProcessNotificationUseCase struct {
	Repo      NotificationRepository
	ProviderA NotificationProvider
}

func NewSaveNotificationUseCase(repo NotificationRepository, providerA NotificationProvider) *ProcessNotificationUseCase {
	return &ProcessNotificationUseCase{Repo: repo, ProviderA: providerA}
}

func (sn *ProcessNotificationUseCase) Execute(ua *entities.UserAction) error {
	notification := initializePendingNotification(ua)

	log.Printf("Getting template for %s.", notification.Type)
	contentWithPlaceholders := sn.Repo.GetTemplateOfNotification(notification.Type)

	finalContent := replaceTemplatePlaceholders(contentWithPlaceholders, notification, ua)
	notification.Content = finalContent

	if err := sn.Repo.Save(notification); err != nil {
		return err
	}

	log.Printf("Sending notification through third-party provider")
	if sn.sendWithRetries(notification, finalContent, sn.ProviderA) {
		return sn.Repo.UpdateNotification(notification)
	}

	log.Printf("The third-party provider haas failed")
	notification.Status = "permanent_failure"
	return sn.Repo.UpdateNotification(notification)
}

func initializePendingNotification(ua *entities.UserAction) *entities.Notification {
	return &entities.Notification{
		ID:           time.Now().Format("20060102150405"),
		Type:         ua.ActionType,
		UserActionID: ua.ID,
		Status:       "pending",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (sn *ProcessNotificationUseCase) sendWithRetries(notification *entities.Notification, finalContent string, provider NotificationProvider) bool {
	retries := 0

	for retries < MaxRetries {
		err := provider.Send(finalContent)
		if err == nil {
			notification.Status = "sent"
			return true
		}

		log.Printf("Error sending notification %s: %v\n", notification.ID, err)

		notification.Status = "retrying"
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

func replaceTemplatePlaceholders(template string, notification *entities.Notification, ua *entities.UserAction) string {
	replacements := map[string]string{
		"{$userID}":           ua.UserID,
		"{$notificationType}": notification.Type,
		"{$amount}":           fmt.Sprintf("%.2f", ua.Amount),
	}

	for placeholder, value := range replacements {
		template = strings.ReplaceAll(template, placeholder, value)
	}

	return template
}
