package models

import "time"

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

type Tansaction struct {
	ID        uint            `gorm:"primaryKey"`
	UserID    uint            `gorm:"not null;index"`
	Amount    float64         `gorm:"type:numeric(15,2);not null"`
	Type      TransactionType `gorm:"type:varchar(10);not null;check:type IN ('income', 'expense')"`
	Timestamp time.Time       `gorm:"type:timestamptz;not null"`
	CreatedAT time.Time       `gorm:"type:timestamptz;default:current_timestamp"`

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
