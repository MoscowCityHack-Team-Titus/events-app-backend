package apiserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/tetovske/events-app-backend/internal/app/repository"
	"github.com/tetovske/events-app-backend/internal/app/repository/implementation"
	"github.com/tetovske/events-app-backend/internal/app/repository/models"
	"gorm.io/gorm"
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

func InitDb(db *sql.DB) {
	resp, err := http.Get(`https://www.mos.ru/api/newsfeed/v4/frontend/json/afisha?expand=spheres&filter={">created_at":"2021-05-15"}`)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}

	events := models.ApiEventsPage{}
	err = json.Unmarshal(body, &events)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, t := range events.Items {
		curEvent := models.Event{
			Model:       gorm.Model{},
			ID: 		 uint(t.ID),
			Title:       t.Title,
			Description: t.Text,
			Sphere: 	 t.Spheres[0].Title,
			Users:       nil,
		}
		_, gormDB, err := repository.NewRepository(db)
		if err != nil {
			fmt.Println(err)
			return
		}
		eventRepo := implementation.NewEventRepo(gormDB)

		eventRepo.CreateEvent(curEvent)
		fmt.Println(t.ID)
	}
}