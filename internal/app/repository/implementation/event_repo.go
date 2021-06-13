package implementation

import (
	"github.com/MCHTitus/events-app-backend/internal/app/repository/models"
	"gorm.io/gorm"
	"log"
)

type EventRepo struct {
	db *gorm.DB
}

func NewEventRepo(db *gorm.DB) *EventRepo {
	err := db.AutoMigrate(&models.Event{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &EventRepo{db: db}
}

func (r *EventRepo) CreateEvent(event models.Event) {
	r.db.Create(&event)
}

func (r *EventRepo) DeleteEvent(event models.Event) {
	r.db.Delete(&event)
}
