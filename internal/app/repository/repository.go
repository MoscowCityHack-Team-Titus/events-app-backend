package repository

import (
	"database/sql"
)

// import "github.com/jmoiron/sqlx"
// import "github.com/golang-migrate/migrate/v4"

type User interface {

}

type Repository struct {
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
