package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tetovske/events-app-backend/config"
	"github.com/tetovske/events-app-backend/internal/app/repository"
	_ "github.com/tetovske/events-app-backend/migrations"
	"log"
)

func main() {
	config.Init()

	db, err := repository.NewPSQLConnection(repository.PSQLConfig{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Password: viper.GetString("db.password"),
		DBName: viper.GetString("db.database"),
		Username: viper.GetString("db.user"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	_, _, err = repository.NewRepository(db) // Второй аргумент для обращения к БД
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("App finished successfully!")
}