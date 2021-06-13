package services

import (
	"github.com/tetovske/events-app-backend/internal/app/repository"
	"github.com/tetovske/events-app-backend/internal/app/repository/models"
)

type RegisterUserJSON struct {
	Gender 	string
	Age 	string
	Hobbies []map[string]string
}

type UserAuthenticationService struct {
	repo *repository.Repository
}

func NewUserAuthenticationService(repo *repository.Repository) *UserAuthenticationService {
	return &UserAuthenticationService{repo: repo}
}

func (r *UserAuthenticationService) ChangeUser(req *RegisterUserJSON) (*models.User, error) {
	for _, hobbie := range req.Hobbies {
		r.repo.GDB.FirstOrCreate(&models.Preference{}, models.Preference{
			Title: hobbie["label"],
		})
	}

	var prefs []*models.Preference
	for _, hobbie := range req.Hobbies {
		var pref models.Preference
		r.repo.GDB.Model(&models.Preference{}).Where("title = ?", hobbie["label"]).First(&pref)
		if &pref != nil {
			prefs = append(prefs, &pref)
		}
	}

	var usr models.User

	r.repo.GDB.First(&usr)
	if &usr != nil {
		usr.Age = req.Age
		usr.Gender = req.Gender
		usr.Preferences = prefs
		r.repo.GDB.Save(&usr)
	} else {
		r.repo.GDB.Create(&models.User{
			Gender:      req.Gender,
			Age:         req.Age,
			Preferences: prefs,
		})
	}

	return &usr, nil
}