package models

import (
	"time"
)

type User struct {
	ID          uint   `gorm:"primarykey"`
	FireBaseUID string `gorm:"unique;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// NickName for show
	NickName string `gorm:"size:64;not null"`
	// SelfIntroduce is introducing themself
	SelfIntroduce string `gorm:"not null"`
	// Sex 0: woman, 1: man, 99: otherwise
	Sex uint `gorm:"not null"`
	// Age: if this is 0 then private
	Age uint `gorm:"not null"`

	Posts []*Post `gorm:"many2many:user_posts;"`

	// Permission 0: normal
	Permission uint

	// Status 0: normal, 1: stranger, 99: danger
	Status uint `gorm:"not null"`

	Initialized bool `gorm:"not null"`
}
