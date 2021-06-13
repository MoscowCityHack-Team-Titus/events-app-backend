package services

import (
	"github.com/tetovske/events-app-backend/internal/app/repository"
	"github.com/tetovske/events-app-backend/internal/app/repository/models"
)

type AddToWishListJSON struct {
	Id 			uint
	Wishlist 	bool
}

type EventService struct {
	repo *repository.Repository
}

func NewEventService(repo *repository.Repository) *EventService {
	return &EventService{repo: repo}
}

func (r *EventService) AddToWishlist(req *AddToWishListJSON) (*models.User, error) {
	var usr models.User

	r.repo.GDB.Take(&usr)
	if &usr != nil {
		var event models.Event
		r.repo.GDB.Model(&models.Event{}).Where("id = ?", req.Id).First(&event)

		if req.Wishlist {
			usr.Events = append(usr.Events, &event)
		} else {
			var updEvents []*models.Event
			for _, event := range usr.Events {
				if event.ID != req.Id {
					updEvents = append(updEvents, event)
				}
			}
			usr.Events = updEvents
		}

		r.repo.GDB.Save(&usr)
		return &usr, nil
	} else {
		return nil, nil
	}
}
