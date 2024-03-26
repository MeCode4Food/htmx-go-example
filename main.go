package main

import (
	"fmt"
	"math/rand"
	"net/http"

	log "github.com/mgutz/logxi/v1"

	"github.com/gorilla/mux"
)

var cats = NewCat(3)

func main() {
	router := mux.NewRouter()
	web := http.Server{
		Addr:    ":7070",
		Handler: router,
	}

	router.HandleFunc("/", GetMainTemplateHandler).Methods(http.MethodGet)
	router.HandleFunc("/cats", GetCatsTemplateHandler).Methods(http.MethodGet)
	router.HandleFunc("/cats/edit/{id}", GetCatEditTemplateHandler).Methods(http.MethodGet)

	router.HandleFunc("/cats/create", GetCatCreateTemplateHandler).Methods(http.MethodGet)

	router.HandleFunc("/cat", PostCreateNewTemplateHandler).Methods(http.MethodPost)

	router.HandleFunc("/cat/{id}", DeleteCatByIDTemplateHandler).Methods(http.MethodDelete)

	router.HandleFunc("/cat/{id}", PutCatByIDTemplateHandler).Methods(http.MethodPut)

	log.Info("starting server", "url", "http://localhost"+web.Addr)
	if err := web.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("error on listen and serve", "error", err)
	}
}

type HxTriggerValue map[string]string

type MainPageContext struct {
	PageTitle string
	Cats      []*Cat
}

var (
	EditTypeCreate = "create"
	EditTypeUpdate = "update"
)

type EditCatPageContext struct {
	EditType  string
	PageTitle string
	Cat       *Cat
}

type CatContext struct {
	Cat *Cat
}

type Cat struct {
	Name  string
	Breed string
	ID    string
}

var breeds = []string{
	"Siamese",
	"Persian",
	"Maine Coon",
	"British Shorthair",
	"Standard Issue Cat",
}

func NewCat(n int) []*Cat {
	cats := make([]*Cat, n)
	for i := 0; i < n; i++ {
		cats[i] = &Cat{
			Name:  "Cat",
			Breed: breeds[rand.Intn(len(breeds))],
			ID:    fmt.Sprintf("%d", rand.Intn(1000000)),
		}
	}
	return cats
}
