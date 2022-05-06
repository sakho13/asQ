package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	// CreatedBy 作成者
	CreatedBy uint
}
