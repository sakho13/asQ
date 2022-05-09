package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	UserID uint
	User   User
	// Password 管理者ログイン時のパスワード
	Password string `gorm:"not null"`
}
