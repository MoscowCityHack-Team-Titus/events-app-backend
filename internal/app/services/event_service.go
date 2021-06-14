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

type SendMessageToChatJSON struct {
	MessageID 	uint
	EventID 	uint
	Message 	string
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

func (r *EventService) GetWishlist() ([]models.ApiSpecificEvent, error) {
	var usr models.User

	r.repo.GDB.Preload("Likes").Find(&usr)
	events := make([]models.ApiSpecificEvent, 0)

	for _, like := range usr.Likes {
		var event models.ApiSpecificEvent

		url := fmt.Sprintf("https://www.mos.ru/api/newsfeed/v4/frontend/json/afisha/%d?expand=spheres", like.EventID)
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		body, _ := ioutil.ReadAll(resp.Body)
		if err = json.Unmarshal(body, &event); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
func (r *EventService) Recommendations() (*models.ApiEventsPage, error) {
	var usr models.User

	r.repo.GDB.Preload("Preferences").Find(&usr)
	events := models.ApiEventsPage{}
	prefPerSection := 15 / len(usr.Preferences)
	for i, pref := range usr.Preferences {
		url := fmt.Sprintf("https://www.mos.ru/api/newsfeed/v4/frontend/json/afisha?expand=spheres&filter={\"spheres.title\":\"%s\"}&per-page=%d", pref.Title, prefPerSection)
		fmt.Println(url)
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
			fmt.Println(string(body))
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

func (r *EventService) SendMessage(req *SendMessageToChatJSON) (*models.Chat, error) {
	var usr	models.User

	r.repo.GDB.First(&usr)

	chat := models.Chat{
		EventID: req.EventID,
	}
	r.repo.GDB.Model(&models.Chat{}).FirstOrCreate(&chat)
	msg := models.Message{
		UserID:  usr.ID,
		Message: req.Message,
	}
	chat.Messages = append(chat.Messages, msg)
	r.repo.GDB.Save(&chat)

	return &chat, nil
}

func (r *EventService) UpdateMessage(req *SendMessageToChatJSON) (*models.Chat, error) {
	var msg models.Chat


	return &msg, nil
}

func (r *EventService) DeleteMessage(req *SendMessageToChatJSON) (*models.Chat, error) {
	var msg models.Chat

	r.repo.GDB.Model(&models.Chat{}).Where("id = ?", req.MessageID).First(&msg)
	r.repo.GDB.Delete(&msg)

	return nil, nil
}
