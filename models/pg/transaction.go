package models

import "time"

type Transaction struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	UserId            uint      `gorm:"not null" json:"user_id"`
	Type              string    `gorm:"size:10;not null" json:"type"` // 'income' or 'expense'
	Amount            float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Date              time.Time `gorm:"not null"`
	IncomeSourceID    *uint     `json:"income_source_id"`    // Foreign key to IncomeSource, nullable
	ExpenseCategoryID *uint     `json:"expense_category_id"` // Foreign key to ExpenseCategory, nullable
	Description       string    `gorm:"type:text" json:"description"`
	CreatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`

	// Associations
	IncomeSource    *IncomeSource    `gorm:"foreignKey:IncomeSourceID;references:ID"`
	ExpenseCategory *ExpenseCategory `gorm:"foreignKey:ExpenseCategoryID;references:ID"`
}
