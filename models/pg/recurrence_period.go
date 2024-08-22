package models

type RecurrencePeriod struct {
	ID                    uint                   `gorm:"primaryKey;autoIncrement" json:"id"`
	PeriodName            string                 `gorm:"size:50;not null" json:"period_name"`
	RecurringTransactions []RecurringTransaction `gorm:"foreignKey:RecurrencePeriodID" json:"recurring_transaction"`
}
