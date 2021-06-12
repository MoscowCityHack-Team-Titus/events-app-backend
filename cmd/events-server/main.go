package main

import (
	"fmt"
	"github.com/MCHTitus/events-app-backend/config"
	"github.com/MCHTitus/events-app-backend/internal/app/apiserver"
	"github.com/MCHTitus/events-app-backend/internal/app/repository"
	_ "github.com/MCHTitus/events-app-backend/migrations"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func main() {
	config.Init()

	time.Sleep(time.Second * 5)

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

	repository.NewRepository(db)
	fmt.Println("App finished successfully!")
	mux := http.NewServeMux()
	mux.HandleFunc("/", apiserver.Home)
	mux.HandleFunc("/events/", apiserver.GetEvents)
	mux.HandleFunc("/event", apiserver.GetEventInfo)
	mux.HandleFunc("/like", apiserver.LikeEvent) // TODO: мейби объединить с '/event', но с методом POST?
	mux.HandleFunc("/chat", apiserver.InteractChat)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
