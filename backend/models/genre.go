package models

import "gorm.io/gorm"

// Genre means classification of Post
//
// write: admin only
type Genre struct {
	gorm.Model

	Title       string `gorm:"not null"`
	Explanation string `gorm:"not null"`
}
