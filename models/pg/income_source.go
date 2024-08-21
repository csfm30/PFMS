package models

import "time"

type IncomeSource struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null"`
	Name        string    `gorm:"size:100;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
