package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID 			uint
	Username 	string
	Email 		string
	Gender 		string
	Age			string
	Events 		[]*Event `gorm:"many2many:events_users;"`
	Preferences []*Preference `gorm:"many2many:preferences_users;"`
	Likes		[]Like
}
