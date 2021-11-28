package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-project/app/domain"

	"github.com/go-chi/chi/v5"
)

type getter interface {
	Index() (*domain.Characters, error)
	Show(characterID int) (*domain.Character, error)
}

type CharactersHandler struct {
	service getter
}


func NewCharactersHandler(getter getter) CharactersHandler{
	return CharactersHandler{getter}
}


func (ch CharactersHandler) Index( w http.ResponseWriter, r *http.Request) {


	characters, err := ch.service.Index()

	if err != nil {
		bytes, _ := json.Marshal(struct {
			Code int `json:"code"`
			Message string `json:"message"`
		}{http.StatusBadRequest, err.Error()})

		w.Header().Add("Content-Type", "application/json")
		w.Write(bytes)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}

func (ch CharactersHandler) Show(w http.ResponseWriter, r *http.Request) {

	characterID := chi.URLParam(r, "id")

	characterIDInt, err := strconv.Atoi(characterID)

	if err != nil {
		bytes, _ := json.Marshal(struct {
			Code int `json:"code"`
			Message string `json:"message"`
		}{http.StatusBadRequest, "id not allowed"})

		w.Header().Add("Content-Type", "application/json")
		w.Write(bytes)
	}

	character, err := ch.service.Show(characterIDInt)

	if err != nil {
		bytes, _ := json.Marshal(struct {
			Code int `json:"code"`
			Message string `json:"message"`
		}{http.StatusBadRequest, err.Error()})

		w.Header().Add("Content-Type", "application/json")
		w.Write(bytes)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(character)
}
