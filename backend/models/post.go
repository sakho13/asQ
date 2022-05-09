package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	// CreatedBy 作成者
	CreatedBy uint
	User      User `gorm:"foreignKey:CreatedBy"`

	Title       string `gorm:"size:128;not null"`
	Explanation string `gorm:"not null"`

	Genres []*Genre `gorm:"many2many:post_genres"`

	Users []*User `gorm:"many2many:user_posts;"`
}
