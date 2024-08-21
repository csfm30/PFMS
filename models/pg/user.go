package models

import "time"

// User model
type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Username     string    `gorm:"size:100;not null"`
	Email        string    `gorm:"size:100;unique;not null"`
	PasswordHash string    `gorm:"size:255;not null"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Role         string    `json:"role"`
	IsSetData    bool      `json:"is_set_data"`
	IsLogin      bool      `json:"is_login"`

	IncomeSources         []IncomeSource         `gorm:"foreignKey:UserID"`
	ExpenseCategories     []ExpenseCategory      `gorm:"foreignKey:UserID"`
	Transactions          []Transaction          `gorm:"foreignKey:UserID"`
	RecurringTransactions []RecurringTransaction `gorm:"foreignKey:UserID"`
}
