package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID 			uint
	Username 	string
	Email 		string
	Events 		[]Event `gorm:"many2many:events_users;"`
}
