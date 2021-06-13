package handlers

import (
	"encoding/json"
	"github.com/tetovske/events-app-backend/internal/app/repository"
	"github.com/tetovske/events-app-backend/internal/app/services"
	"log"
	"net/http"
)

type UserHandler struct {
	repo *repository.Repository
}

func NewUserHandler(repo *repository.Repository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (rc *UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req services.RegisterUserJSON

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println(err)
			http.Error(w, "Invalid request", 422)
			return
		}

		usr, err := services.NewService(rc.repo).Authentication.ChangeUser(&req)
		if err != nil {
			http.Error(w, "Error on handling user auth", 500)
			return
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

func (rc *UserHandler) Wishlist(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req services.AddToWishListJSON

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println(err)
			http.Error(w, "Invalid request", 422)
			return
		}

		usr, err := services.NewService(rc.repo).EventManager.AddToWishlist(&req)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		resp, _ := json.Marshal(usr)
		if _, err := w.Write(resp); err != nil {
			log.Fatal(err)
		}
	} else if r.Method == http.MethodGet {
		events, err := services.NewService(rc.repo).EventManager.GetWishlist()
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		resp, _ := json.Marshal(events)
		if _, err := w.Write(resp); err != nil {
			log.Fatal(err)
		}
	} else {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод не дозволен", 405)
		return
	}
}

func (rc *UserHandler) Recommendations(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		events, err := services.NewService(rc.repo).EventManager.Recommendations()
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		resp, _ := json.Marshal(events)
		if _, err := w.Write(resp); err != nil {
			log.Fatal(err)
		}
	} else {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод не дозволен", 405)
		return
	}
}
