package repositories

import (
	"github.com/rodrinoblega/notification_handler/src/entities"
	"gorm.io/gorm"
	"log"
	"time"
)

const DefaultWording = "You have one new notification, check your application to see more information"

type PostgresNotificationRepository struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *PostgresNotificationRepository {
	return &PostgresNotificationRepository{DB: db}
}

func (repo *PostgresNotificationRepository) Save(n *entities.Notification) error {
	return repo.DB.Create(n).Error
}

func (repo *PostgresNotificationRepository) UpdateStatus(id string, status string) (*entities.Notification, error) {
	notification := &entities.Notification{}

	err := repo.DB.Model(notification).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":     status,
			"updated_at": time.Now(),
		}).Error
	if err != nil {
		return nil, err
	}

	if err := repo.DB.First(notification, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return notification, nil
}

func (repo *PostgresNotificationRepository) UpdateRetries(id string, retries int) (*entities.Notification, error) {
	notification := &entities.Notification{}

	err := repo.DB.Model(notification).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"retries":    retries,
			"updated_at": time.Now(),
		}).Error
	if err != nil {
		return nil, err
	}

	if err := repo.DB.First(notification, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return notification, nil
}

func (repo *PostgresNotificationRepository) UpdateNotification(n *entities.Notification) error {
	return repo.DB.Model(n).
		Select("Status", "Retries", "UpdatedAt", "Content").
		Updates(map[string]interface{}{
			"Status":    n.Status,
			"Retries":   n.Retries,
			"UpdatedAt": time.Now(),
			"Content":   n.Content,
		}).Error
}

func (repo *PostgresNotificationRepository) GetFailedNotifications() ([]entities.Notification, error) {
	var notifications []entities.Notification
	err := repo.DB.Where("status IN ?", []string{"permanent_failure", "permanent_failure_async"}).Find(&notifications).Error
	return notifications, err
}

func (repo *PostgresNotificationRepository) GetNotifications(userID string, notificationType string, startDate, endDate time.Time, limit, offset int) ([]entities.Notification, error) {
	var notifications []entities.Notification

	query := repo.DB.
		Joins("JOIN user_actions ON user_actions.id = notifications.user_action_id").
		Where("user_actions.user_id = ?", userID)

	if notificationType != "" {
		query = query.Where("notifications.type = ?", notificationType)
	}
	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("notifications.created_at BETWEEN ? AND ?", startDate, endDate)
	}

	err := query.
		Order("notifications.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&notifications).Error
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (repo *PostgresNotificationRepository) GetTemplateOfNotification(notificationType string) string {
	var template entities.NotificationTemplate

	if err := repo.DB.Where("type = ?", notificationType).First(&template).Error; err != nil {
		log.Printf("no template found for notification type: %s", notificationType)
	}

	if template.Template == "" {
		return DefaultWording
	}

	return template.Template

}
