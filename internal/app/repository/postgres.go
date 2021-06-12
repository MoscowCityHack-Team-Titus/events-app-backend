package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"log"
)

type PSQLConfig struct {
	Host string
	Port string
	Username string
	Password string
	DBName string
	SSLMode string
}

const driverName = "postgres"

func NewPSQLConnection(cnf PSQLConfig) (*sql.DB, error) {
	db, err := sql.Open(driverName, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cnf.Host, cnf.Port, cnf.Username, cnf.DBName, cnf.Password, cnf.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println(fmt.Sprintf("Connected to db: %s", cnf.DBName))

	log.Println("Starting migration...")
	err = goose.Up(db, "migrations")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migration Completed")

	return db, nil
}
