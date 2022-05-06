package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// Status 0: normal, 1: stranger, 99: danger
	Status uint
}
