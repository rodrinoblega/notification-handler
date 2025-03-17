package repositories

import (
	"github.com/rodrinoblega/notification_handler/src/entities"
	"gorm.io/gorm"
	"time"
)

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
		Select("Status", "Retries", "UpdatedAt").
		Updates(map[string]interface{}{
			"Status":    n.Status,
			"Retries":   n.Retries,
			"UpdatedAt": time.Now(),
		}).Error
}

func (repo *PostgresNotificationRepository) GetFailedNotifications() ([]entities.Notification, error) {
	var notifications []entities.Notification
	err := repo.DB.Where("status IN ?", []string{"permanent_failure", "permanent_failure_async"}).Find(&notifications).Error
	return notifications, err
}
