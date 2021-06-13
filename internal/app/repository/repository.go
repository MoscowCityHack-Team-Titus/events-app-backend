package repository

import (
	"database/sql"
	"github.com/tetovske/events-app-backend/internal/app/repository/implementation"
	"github.com/tetovske/events-app-backend/internal/app/repository/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User interface {
	CreateUser(user models.User)
	DeleteUser(user models.User)
}

type Event interface {
	CreateEvent(event models.Event)
	DeleteEvent(event models.Event)
}

type Preference interface {
	CreatePreference(pref models.Preference)
	DeletePreference(pref models.Preference)
}

type Repository struct {
	GDB *gorm.DB
	User
	Event
	Preference
}

func NewRepository(db *sql.DB) (*Repository, *gorm.DB, error) {
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, gdb, err
	}

	return &Repository{
		GDB: gdb,
		User: implementation.NewUserRepo(gdb),
		Event: implementation.NewEventRepo(gdb),
		Preference: implementation.NewPreferenceRepo(gdb),
	}, gdb, nil
}
