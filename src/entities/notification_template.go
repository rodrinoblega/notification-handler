package entities

type Template struct {
	TemplateName string `gorm:"type:text;primaryKey;not null"`
	Content      string `gorm:"type:text;not null"`
}

type NotificationTemplate struct {
	Type     string `gorm:"primaryKey" json:"type"`
	Template string `gorm:"not null" json:"template"`
}
