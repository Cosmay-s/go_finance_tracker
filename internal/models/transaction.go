package models

import (
	"errors"
	"time"
)

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

type Transaction struct {
	ID        uint            `gorm:"primaryKey"`
	UserID    uint            `gorm:"not null;index"`
	Amount    float64         `gorm:"type:numeric(15,2);not null"`
	Type      TransactionType `gorm:"type:varchar(10);not null;check:type IN ('income', 'expense')"`
	Timestamp time.Time       `gorm:"type:timestamptz;not null"`
	CreatedAt time.Time       `gorm:"type:timestamptz;default:current_timestamp"`

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("сумма транзакции должна быть положительной")
	}

	if t.Type != Income && t.Type != Expense {
		return errors.New("тип транзакции должен быть 'income' или 'expense'")
	}

	if t.Timestamp.IsZero() {
		return errors.New("время транзакции обязательно")
	}

	if t.UserID == 0 {
		return errors.New("транзакция должна быть привязана к пользователю")
	}

	return nil
}
