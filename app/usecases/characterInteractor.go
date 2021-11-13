package usecases

import (
	"go-project/app/domain"
)

// Return all characters
func Index() (characters domain.Characters) {
	characters  = FindAll()

	return
}

// Return character by ID
func Show(characterID int) (character domain.Characters){
	character = FindByID(characterID)

	return
}

