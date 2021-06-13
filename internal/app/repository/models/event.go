package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	ID 			uint
	Title 		string
	Description string
	Sphere		string
	Users 		[]*User `gorm:"many2many:events_users;"`
	MapLng		float64
	MapLat		float64
}
