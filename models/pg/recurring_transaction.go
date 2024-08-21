package models

import "time"

type RecurringTransaction struct {
	ID                 uint       `gorm:"primaryKey;autoIncrement"`
	UserID             uint       `gorm:"not null"`
	Type               string     `gorm:"size:10;not null"` // 'income' or 'expense'
	Amount             float64    `gorm:"type:decimal(10,2);not null"`
	StartDate          time.Time  `gorm:"not null"`
	EndDate            *time.Time // Nullable
	RecurrencePeriodID uint       `gorm:"not null"`
	SourceID           *uint      // Foreign key to IncomeSource or ExpenseCategory
	Description        string     `gorm:"type:text"`
	CreatedAt          time.Time  `gorm:"default:CURRENT_TIMESTAMP"`

	// Associations
	RecurrencePeriod RecurrencePeriod `gorm:"foreignKey:RecurrencePeriodID;references:ID"`
	IncomeSource     *IncomeSource    `gorm:"foreignKey:SourceID;references:ID"`
	ExpenseCategory  *ExpenseCategory `gorm:"foreignKey:SourceID;references:ID"`
}
