package models

type Saving struct {
	ID              uint    `gorm:"primaryKey;autoIncrement"`
	UserId          uint    `gorm:"not null" json:"user_id"`
	Name            string  `gorm:"not null" json:"name"`
	TargetAmount    float64 `gorm:"type:decimal(10,2);not null" json:"target_amount"`
	CurrentSaving   float64 `gorm:"type:decimal(10,2);not null" json:"current_saving"`
	AmountSaved     float64 `gorm:"type:decimal(10,2);not null" json:"amount_saved"`
	RemainingAmount float64 `gorm:"type:decimal(10,2);not null" json:"remaining_amount"`
}
