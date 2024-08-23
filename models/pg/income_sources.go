package models

import "time"

type IncomeSource struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId      uint      `gorm:"not null" json:"user_id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
