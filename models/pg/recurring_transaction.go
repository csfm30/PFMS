package models

import "time"

type RecurringTransaction struct {
	ID                 uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId             uint       `gorm:"not null" json:"user_id"`
	Type               string     `gorm:"size:10;not null" json:"type"` // 'income' or 'expense'
	Amount             float64    `gorm:"type:decimal(10,2);not null" json:"amount"`
	StartDate          time.Time  `gorm:"not null" json:"start_date"`
	EndDate            *time.Time // Nullable`json:"end_date"`
	RecurrencePeriodID uint       `gorm:"not null" json:"recurren_period_id"`
	SourceID           *uint      // Foreign key to IncomeSource or ExpenseCategory`json:"source_id"`
	Description        string     `gorm:"type:text" json:"description"`
	CreatedAt          time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`

	// Associations
	RecurrencePeriod RecurrencePeriod `gorm:"foreignKey:RecurrencePeriodID;references:ID"`
	IncomeSource     *IncomeSource    `gorm:"foreignKey:SourceID;references:ID"`
	ExpenseCategory  *ExpenseCategory `gorm:"foreignKey:SourceID;references:ID"`
}
