package handlers

import (
	"encoding/json"
	"github.com/tetovske/events-app-backend/internal/app/repository"
	"github.com/tetovske/events-app-backend/internal/app/services"
	"log"
	"net/http"
)

type EventHandler struct {
	repo *repository.Repository
}

func NewEventHandler(repo *repository.Repository) *EventHandler {
	return &EventHandler{repo: repo}
}

func (rc *EventHandler) SendToChat(w http.ResponseWriter, r *http.Request) {
	var req services.SendMessageToChatJSON

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
		http.Error(w, "Invalid request", 422)
		return
	}

	if r.Method == http.MethodPost {
		chat, err := services.NewService(rc.repo).EventManager.SendMessage(&req)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		resp, _ := json.Marshal(chat)
		if _, err := w.Write(resp); err != nil {
			log.Fatal(err)
		}
	} else if r.Method == http.MethodGet {
		usr, err := services.NewService(rc.repo).EventManager.GetWishlist()
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		resp, _ := json.Marshal(usr)
		if _, err := w.Write(resp); err != nil {
			log.Fatal(err)
		}
	} else {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод не дозволен", 405)
		return
	}
}
