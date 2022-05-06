package models

import "time"

type Admin struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UserID    uint
	User      User
	Password  string
}
