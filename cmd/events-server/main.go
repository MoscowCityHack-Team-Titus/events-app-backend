package main

import (
	"github.com/brshpl/events-app-backend/internal/app/apiserver"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", apiserver.Home)
	mux.HandleFunc("/events/", apiserver.GetEvents)
	mux.HandleFunc("/event", apiserver.GetEventInfo)
	mux.HandleFunc("/like", apiserver.LikeEvent) // TODO: мейби объединить с '/event', но с методом POST?
	mux.HandleFunc("/chat", apiserver.InteractChat)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
