package models

import "gorm.io/gorm"

type Preference struct {
	gorm.Model
	ID 		uint
	Title 	string
	Users 	[]*User `gorm:"many2many:preferences_users;"`
}
