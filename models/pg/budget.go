package models

import "time"

type Budget struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	UserId     uint      `gorm:"not null"`
	CategoryID uint      `gorm:"not null"` // ForeignKey to ExpenseCategory
	Amount     float64   `gorm:"type:decimal(10,2);not null"`
	StartDate  time.Time `gorm:"not null"`
	EndDate    time.Time `gorm:"not null"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
