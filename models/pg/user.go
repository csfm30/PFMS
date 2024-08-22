package models

import "time"

// User model
type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Username     string    `gorm:"size:100;not null" json:"username"`
	Email        string    `gorm:"size:100;unique;not null" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"password_hash"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Role         string    `json:"role"`
	IsSetData    bool      `json:"is_set_data"`
	IsLogin      bool      `json:"is_login"`

	IncomeSources         []IncomeSource         `gorm:"foreignKey:UserId"`
	ExpenseCategories     []ExpenseCategory      `gorm:"foreignKey:UserId"`
	Transactions          []Transaction          `gorm:"foreignKey:UserId"`
	RecurringTransactions []RecurringTransaction `gorm:"foreignKey:UserId"`
}
