package usecases

import (
	"github.com/rodrinoblega/notification_handler/src/entities"
)

type EventPublisher interface {
	Publish(userAction *entities.UserAction) error
}

type UserActionRepository interface {
	Create(userAction *entities.UserAction) error
	GetByID(id int) (*entities.UserAction, error)
}

type PublishEventUseCase struct {
	Publisher            EventPublisher
	UserActionRepository UserActionRepository
}

func NewPublishEventUseCase(publisher EventPublisher, repo UserActionRepository) *PublishEventUseCase {
	return &PublishEventUseCase{Publisher: publisher, UserActionRepository: repo}
}

func (uc *PublishEventUseCase) SendEvent(ua *entities.UserAction) error {
	err := uc.UserActionRepository.Create(ua)
	if err != nil {
		return err
	}
	return uc.Publisher.Publish(ua)
}
