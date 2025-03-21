package repositories

import (
	"github.com/rodrinoblega/notification_handler/src/entities"

	"gorm.io/gorm"
)

type PostgresTemplateRepository struct {
	DB *gorm.DB
}

func NewPostgresTemplateRepository(DB *gorm.DB) *PostgresTemplateRepository {
	return &PostgresTemplateRepository{DB: DB}
}

type TemplateResponse struct {
	Type         string `json:"type"`
	TemplateName string `json:"template_name"`
	Content      string `json:"content"`
}

func (ptr *PostgresTemplateRepository) CreateTemplate(template *entities.NotificationTemplate) error {
	return ptr.DB.Create(template).Error
}

func (ptr *PostgresTemplateRepository) GetAllTemplates() ([]entities.NotificationTemplate, error) {
	var templates []entities.NotificationTemplate
	if err := ptr.DB.Find(&templates).Error; err != nil {
		return nil, err
	}
	return templates, nil
}

func (ptr *PostgresTemplateRepository) UpdateTemplate(notificationType, newTemplate string) error {
	result := ptr.DB.Model(&entities.NotificationTemplate{}).
		Where("type = ?", notificationType).
		Update("template", newTemplate)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (ptr *PostgresTemplateRepository) DeleteTemplate(notificationType string) error {
	result := ptr.DB.Where("type = ?", notificationType).Delete(&entities.NotificationTemplate{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
