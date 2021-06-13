package services

import (
	"github.com/tetovske/events-app-backend/internal/app/repository"
	"github.com/tetovske/events-app-backend/internal/app/repository/models"
)

type Authentication interface {
	ChangeUser(req *RegisterUserJSON) (*models.User, error)
}

type EventManager interface {
	AddToWishlist(req *AddToWishListJSON) (*models.User, error)
	GetWishlist() (*models.User, error)
	Recommendations() (*models.ApiEventsPage, error)
}

type Service struct {
	Authentication
	EventManager
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authentication: NewUserAuthenticationService(repo),
		EventManager: NewEventService(repo),
	}
}
