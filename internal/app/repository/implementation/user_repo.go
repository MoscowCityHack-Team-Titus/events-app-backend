package implementation

import (
	"github.com/MCHTitus/events-app-backend/internal/app/repository/models"
	"gorm.io/gorm"
	"log"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user models.User) {
	r.db.Create(&user)
}

func (r *UserRepo) DeleteUser(user models.User) {
	r.db.Delete(&user)
}