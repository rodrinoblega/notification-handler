package entities

import "time"

type Notification struct {
	ID           string    `gorm:"type:text;primaryKey;not null" json:"id"`
	UserActionID int       `gorm:"not null;index" json:"user_action_id"`
	Type         string    `gorm:"type:text;not null" json:"type"`
	Status       string    `gorm:"type:text;not null" json:"status"`
	Content      string    `gorm:"type:text;not null" json:"content"`
	Retries      int       `gorm:"default:0;not null" json:"retries"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
