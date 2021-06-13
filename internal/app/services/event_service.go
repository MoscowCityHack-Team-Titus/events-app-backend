package services

import (
	"encoding/json"
	"fmt"
	"github.com/tetovske/events-app-backend/internal/app/repository"
	"github.com/tetovske/events-app-backend/internal/app/repository/models"
	"io/ioutil"
	"net/http"
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

func (r *EventService) GetWishlist() (*models.User, error) {
	var usr models.User

	r.repo.GDB.Preload("Likes").Find(&usr)

	return &usr, nil
}
func (r *EventService) Recommendations() (*models.ApiEventsPage, error) {
	var usr models.User

	r.repo.GDB.Preload("Preferences").Find(&usr)
	events := models.ApiEventsPage{}
	prefPerSection := 10 / len(usr.Preferences)
	for i, pref := range usr.Preferences {
		url := fmt.Sprintf("https://www.mos.ru/api/newsfeed/v4/frontend/json/afisha?expand=spheres&filter={\"spheres.title\":\"%s\"}&per-page=%d", pref.Title, prefPerSection)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		body, _ := ioutil.ReadAll(resp.Body)
		if i == 0 {
			if err = json.Unmarshal(body, &events); err != nil {
				return nil, err
			}
		} else {
			temp := models.ApiEventsPage{}
			if err = json.Unmarshal(body, &temp); err != nil {
				return nil, err
			}

			for _, item := range temp.Items {
				events.Items = append(events.Items, item)
			}
		}
	}

	return &events, nil
}
