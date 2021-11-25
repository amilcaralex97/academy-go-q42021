package usecases

import (
	"errors"
	"go-project/app/domain"
)

// Return all characters
func Index() (characters *domain.Characters, err error) {
	characters, err  = FindAll()

	if err != nil {
		return nil, errors.New("error while trying to obtain characters")
	}

	return
}

// Return character by ID
func Show(characterID int) (character *domain.Character, err error){
	character, err = FindByID(characterID)

		if err != nil {
		return nil, errors.New("error while trying to obtain character")
	}

	return
}

