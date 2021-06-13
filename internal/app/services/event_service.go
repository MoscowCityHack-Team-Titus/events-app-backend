package services

import (
	"github.com/tetovske/events-app-backend/internal/app/repository"
	"github.com/tetovske/events-app-backend/internal/app/repository/models"
)

type AddToWishListJSON struct {
	Id 			int
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

	r.repo.GDB.Preload("Likes").Find(&usr)

	if req.Wishlist {
		var likesCount int64

		r.repo.GDB.Model(&models.Like{}).Where("user_id = ? and event_id = ?", usr.ID, req.Id).Count(&likesCount)

		if likesCount < 1 {
			usr.Likes = append(usr.Likes, models.Like{EventID: req.Id})
			r.repo.GDB.Save(&usr)
		}
	} else {
		for _, like := range usr.Likes {
			if like.EventID == req.Id {
				r.repo.GDB.Delete(&like)
				break
			}
		}
		r.repo.GDB.Preload("Likes").Find(&usr)
	}
	return &usr, nil
}

func (r *EventService) GetWishlist(req *AddToWishListJSON) (*models.User, error) {
	var usr models.User

	r.repo.GDB.Preload("Likes").Find(&usr)

	return &usr, nil
}
