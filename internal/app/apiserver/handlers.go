package apiserver

import (
	"fmt"
	"net/http"
	"strconv"
)

// TODO: разобраться, как идентифицировать пользователя (JWT, Cookies or smth)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// TODO: куча работы, связанной с показом релевантных событий

	_, err := w.Write([]byte("Привет из Афиши"))
	if err != nil {
		return
	}
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	respByteString, err := getEventsFromApi(r)
	if err != nil {
		http.Error(w, "400 Ошибка в запросе", http.StatusBadRequest)
		return
	}

	_, err = w.Write(respByteString)
	if err != nil {
		http.Error(w, "400 Ошибка в запросе", http.StatusBadRequest)
		return
	}
}

func GetEventInfo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	_, err = fmt.Fprintf(w, "Количество лайков, меток, всякой всячины для ивента %d...", id)
	if err != nil {
		return
	}
}

func LikeEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод не дозволен", 405)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// TODO: Проверяем в БД, лайкал ли этот пользователь это мероприятие - выставляем лайк, если нет

	_, err = w.Write([]byte("Лайкнул событие? - Красавчик!"))
	if err != nil {
		return
	}
}

func InteractChat(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodPost {
		// TODO: Постим сообщение в чятик

		_, err = w.Write([]byte("Поболтал о событии? - Красавчик!"))
		if err != nil {
			return
		}
		return
	} else if r.Method == http.MethodGet {
		// TODO: выдаём последние сообщения в чатике

		_, err = fmt.Fprintf(w, "Вот тебе ламповый чятик для ивента %d...", id)
		if err != nil {
			return
		}
	} else {
		w.Header().Set("Allow", http.MethodPost) // TODO: разобраться, можно ли разрешить несколько методов
		http.Error(w, "Метод не дозволен", 405)
		return
	}
}
