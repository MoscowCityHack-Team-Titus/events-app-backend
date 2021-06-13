package implementation

import (
	"github.com/tetovske/events-app-backend/internal/app/repository/models"
	"gorm.io/gorm"
	"log"
)

type PreferenceRepo struct {
	db *gorm.DB
}

func NewPreferenceRepo(db *gorm.DB) *PreferenceRepo {
	err := db.AutoMigrate(&models.Preference{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &PreferenceRepo{db: db}
}

func (r *PreferenceRepo) CreatePreference(pref models.Preference) {
	r.db.Create(&pref)
}

func (r *PreferenceRepo) DeletePreference(pref models.Preference) {
	r.db.Delete(&pref)
}
