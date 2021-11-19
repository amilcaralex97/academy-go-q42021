package infrastructure

import (
	"log"
	"net/http"

	"go-project/app/interfaces"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Dispatch is handle routing
func Dispatch() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	
	r.Route("/characters", func(r chi.Router){
		r.Get("/", interfaces.Index)
		r.Get("/{id}", interfaces.Show)
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Printf("%s", err)
	}
}