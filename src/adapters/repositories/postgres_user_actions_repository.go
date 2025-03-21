package repositories

import (
	"github.com/rodrinoblega/notification_handler/src/entities"

	"gorm.io/gorm"
)

type PostgresUserActionRepository struct {
	DB *gorm.DB
}

func NewUserActionRepository(DB *gorm.DB) *PostgresUserActionRepository {
	return &PostgresUserActionRepository{DB: DB}
}

func (uar *PostgresUserActionRepository) Create(userAction *entities.UserAction) error {
	return uar.DB.Create(userAction).Error
}

func (uar *PostgresUserActionRepository) GetByID(id int) (*entities.UserAction, error) {
	var action entities.UserAction
	err := uar.DB.Where("id = ?", id).First(&action).Error
	if err != nil {
		return nil, err
	}
	return &action, nil
}
