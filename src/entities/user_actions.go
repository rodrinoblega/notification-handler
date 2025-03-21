package entities

import "time"

type UserAction struct {
	ID         int       `gorm:"primaryKey;type:TEXT" json:"id"`
	UserID     string    `gorm:"type:TEXT;index" json:"user_id"`
	ActionType string    `gorm:"type:TEXT;not null" json:"action_type"`
	Amount     float64   `gorm:"type:numeric(20,6);not null" json:"amount"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}
