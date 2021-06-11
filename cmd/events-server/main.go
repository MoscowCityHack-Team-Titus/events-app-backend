package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/event", getEventInfo)
	mux.HandleFunc("/like", likeEvent) // TODO: мейби объединить с '/event', но с методом POST?
	mux.HandleFunc("/chat", interactChat)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
