package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-project/app/usecases"
)

func Index( w http.ResponseWriter, r *http.Request) {


	characters := usecases.Index()

	// if err != nil  {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(500)
	// 	json.NewEncoder(w).Encode(err)
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}

func Show(w http.ResponseWriter, r *http.Request) {

	characterID, _ := strconv.Atoi(r.URL.Query().Get("id"))



	character := usecases.Show(characterID)
	// if err != nil {

	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(500)
	// 	json.NewEncoder(w).Encode(err)
	// }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(character)
}
