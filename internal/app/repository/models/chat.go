package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	ID 			uint
	EventID		uint
	Messages	[]Message
}

type Message struct {
	gorm.Model
	ID 			uint
	UserID 		uint
	ChatID		uint
	Message		string
}