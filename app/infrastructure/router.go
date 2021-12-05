package infrastructure

import (
	"log"
	"net/http"

	"go-project/app/interfaces"
	"go-project/app/repository"
	"go-project/app/usecases"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const CSVFile string = "/Users/alejandrosanchez/Documents/go_bootcamp/app/resources/characters.csv"
const URL string = "https://swapi.dev/api/people"

// Dispatch is handle routing
func Dispatch() {
	//dependendy injection
	charactersService := usecases.NewCharactersInteractor(repository.NewCharacterRepo(), repository.NewApiRepo(URL), repository.NewCsvRepo(CSVFile))
	charactersHandler := interfaces.NewCharactersHandler(charactersService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/characters", func(r chi.Router) {
		r.Get("/", charactersHandler.Index)
		r.Get("/{id}", charactersHandler.Show)
		r.Get("/fetch-characters", charactersHandler.FetchCharacters)
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Printf("%s", err)
	}
}
