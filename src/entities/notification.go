package entities

import "time"

type Notification struct {
	ID        string    `gorm:"type:text;primaryKey;not null"`
	UserID    string    `gorm:"type:text;not null"`
	Content   string    `gorm:"type:text;not null"`
	Status    string    `gorm:"type:text;not null"`
	Retries   int       `gorm:"default:0;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
