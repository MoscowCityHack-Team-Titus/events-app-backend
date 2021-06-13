package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/tetovske/events-app-backend/internal/app/repository"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetEventsFromApi(r *http.Request) ([]byte, error) {
	getOnlyApiParams(r)

	fmt.Printf("\n\n")
	fmt.Println(r.RequestURI)

	u, err := url.Parse("https://www.mos.ru/api/newsfeed/v4/frontend/json")
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	curUrl := strings.TrimPrefix(r.RequestURI, "/events")
	rel := u.String() + curUrl

	resp, err := http.Get(rel)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return nil, err
	}

	var events []repository.Event

	err = json.Unmarshal(body, &events)
	if err != nil {
		log.Println("Error Unmarshal")
		println("Error Unmarshal")
	}

	return body, nil
}

func getOnlyApiParams(r *http.Request) *http.Request {
	// TODO: не работает и непонятно почему. Я спать
	r.URL.Query().Del("age")
	r.URL.Query().Del("hobbies")
	r.URL.Query().Del("gender")
	fmt.Printf("\n\n")
	fmt.Println(r.RequestURI)
	return r
}