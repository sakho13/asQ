package models

import (
	"time"
)

// Diary means the memory of a Post
type Diary struct {
	PostID uint `gorm:"primarykey"`
	Post   Post `gorm:"foreignKey:PostID"`

	// CreatedBy 作成者
	CreatedBy uint `gorm:"primarykey"`
	User      User `gorm:"foreignKey:CreatedBy"`

	OrderNo uint `gorm:"primarykey"`

	CreatedAt time.Time
	UpdatedAt time.Time

	// Date was the user specified
	Date time.Time `gorm:"not null"`

	Body string `gorm:"not null"`

	IsEdited bool `gorm:"not null"`
}
