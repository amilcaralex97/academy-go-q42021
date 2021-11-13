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
	
	r.Get("/characters", interfaces.Index)
	r.Get("/characters/:id", interfaces.Show)

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Printf("%s", err)
	}
}