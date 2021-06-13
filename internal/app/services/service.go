package services

import (
	"github.com/tetovske/events-app-backend/internal/app/repository"
	"github.com/tetovske/events-app-backend/internal/app/repository/models"
)

type Authentication interface {
	ChangeUser(req *RegisterUserJSON) (*models.User, error)
}

type Service struct {
	Authentication
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Authentication: NewUserAuthenticationService(repo)}
}
