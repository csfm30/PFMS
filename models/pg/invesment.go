package models

type Investment struct {
	ID                  uint    `gorm:"primaryKey;autoIncrement"`
	UserId              uint    `gorm:"not null" json:"user_id"`
	Name                string  `gorm:"not null" json:"name"`
	InitialAmount       float64 `gorm:"type:decimal(10,2);not null" json:"initial_amount"`
	CurrentValue        float64 `gorm:"type:decimal(10,2);not null" json:"current_value"`
	MonthlyContribution float64 `gorm:"type:decimal(10,2);not null" json:"monthly_contribution"`
	Return              float64 `gorm:"type:decimal(10,2);not null" json:"return"`
	Notes               string  `gorm:"size:100;not null" json:"notes"`
}
