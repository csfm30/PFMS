package models

type RecurrencePeriod struct {
	ID                    uint                   `gorm:"primaryKey;autoIncrement"`
	PeriodName            string                 `gorm:"size:50;not null"`
	RecurringTransactions []RecurringTransaction `gorm:"foreignKey:RecurrencePeriodID"`
}
