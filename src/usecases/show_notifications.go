package usecases

import (
	"github.com/rodrinoblega/notification_handler/src/entities"
	"time"
)

type ShowNotificationsUseCase struct {
	Repo NotificationRepository
}

func NewShowNotificationsUseCase(repo NotificationRepository) *ShowNotificationsUseCase {
	return &ShowNotificationsUseCase{Repo: repo}
}

func (sn *ShowNotificationsUseCase) Search(userID string, notificationType string, startDate, endDate time.Time, limit, offset int) ([]entities.Notification, error) {
	notifications, err := sn.Repo.GetNotifications(userID, notificationType, startDate, endDate, limit, offset)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}
