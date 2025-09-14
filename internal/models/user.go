package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Email        string    `gorm:"unique;not null;type:varchar(255)"`
	PasswordHash string    `gorm:"not null;type:varchar(255)"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
