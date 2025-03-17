package usecases

import (
	"github.com/rodrinoblega/notification_handler/src/entities"
)

type EventPublisher interface {
	Publish(notification entities.Notification) error
}

type PublishEventUseCase struct {
	Publisher EventPublisher
}

func NewPublishEventUseCase(publisher EventPublisher) *PublishEventUseCase {
	return &PublishEventUseCase{Publisher: publisher}
}

func (uc *PublishEventUseCase) SendNotification(n entities.Notification) error {
	return uc.Publisher.Publish(n)
}
