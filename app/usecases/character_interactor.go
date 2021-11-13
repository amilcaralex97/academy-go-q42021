package usecases

import (
	"go-project/app/domain"
)

type CharacterInteractor struct {
	CharacterRepository CharacterRepository
}

// Return all characters
func (ci *CharacterInteractor) Index() (characters domain.Characters, err error) {
	characters, err = ci.CharacterRepository.FindAll()

	return
}

// Return character by ID
func (ci *CharacterInteractor) Show(characterID int) (character domain.Character, err error){
	character, err = ci.CharacterRepository.FindByID(characterID)

	return
}

