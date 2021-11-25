package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-project/app/usecases"

	"github.com/go-chi/chi/v5"
)

func Index( w http.ResponseWriter, r *http.Request) {


	characters, err := usecases.Index()

	if err != nil  {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}

func Show(w http.ResponseWriter, r *http.Request) {

	characterID := chi.URLParam(r, "id")

	characterIDInt, err := strconv.Atoi(characterID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

	character, err := usecases.Show(characterIDInt)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(character)
}
