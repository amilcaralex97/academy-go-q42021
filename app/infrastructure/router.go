package infrastructure

import (
	"log"
	"net/http"

	"go-project/app/interfaces"
	"go-project/app/usecases"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Dispatch is handle routing
func Dispatch() {
	//dependendy injection
	charactersService := usecases.NewCharactersInteractor(usecases.NewCharacterRepo())
	charactersHandler := interfaces.NewCharactersHandler(charactersService)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/characters", func(r chi.Router){
		r.Get("/", charactersHandler.Index)
		r.Get("/{id}", charactersHandler.Show)
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Printf("%s", err)
	}
}