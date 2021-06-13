package models

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	ID 			uint
	UserID 		uint
	EventID		int
}
