package main

import (
	"github.com/spf13/viper"
	"github.com/tetovske/events-app-backend/config"
	"github.com/tetovske/events-app-backend/internal/app/apiserver/handlers"
	"github.com/tetovske/events-app-backend/internal/app/repository"
	_ "github.com/tetovske/events-app-backend/migrations"
	"log"
	"net/http"
)

// @title Events App API
// @version 1.0
// @description Планирование досуговых мероприятий в приложении "Моя Москва". Серверная часть.

// @host localhost:4000
// @BasePath /
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

	repo, _, err := repository.NewRepository(db)
	if err != nil {
		log.Fatal(err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/events/", handlers.GetEvents)
	mux.HandleFunc("/event", handlers.GetEventInfo)
	mux.HandleFunc("/like", handlers.LikeEvent)
	mux.HandleFunc("/chat", handlers.InteractChat)

	userHandler := handlers.NewUserHandler(repo)
	mux.HandleFunc("/register", userHandler.RegisterUserHandler)
	mux.HandleFunc("/wishlist", userHandler.Wishlist)
	mux.HandleFunc("/recommendations", userHandler.Recommendations)

	eventHandler := handlers.NewEventHandler(repo)
	mux.HandleFunc("/comment", eventHandler.SendToChat)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	if err = http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal(err)
	}
}
