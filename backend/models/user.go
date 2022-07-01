package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"primarykey"`
	FireBaseUID string    `gorm:"unique;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// NickName for show
	NickName string `gorm:"size:64;not null"`

	TwitterID string `gorm:"not null"`
}
