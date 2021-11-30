package usecases

import (
	"errors"

	"go-project/app/domain"
)

type getter interface {
	FindAll() (*domain.Characters, error)
	FindByID(characterID int) (*domain.Character, error)
	FetchCharacters() (*domain.Characters, error)
}

type CharactersInteractor struct {
	repo getter
}

func NewCharactersInteractor(repo CharactersRepo) CharactersInteractor {
	return CharactersInteractor{repo}
}

//Return fetched characters
func (ci CharactersInteractor) FetchCharacters() (characters *domain.Characters, err error) {
	characters, err = ci.repo.FetchCharacters()

	return
}

// Return all characters
func (ci CharactersInteractor) Index() (characters *domain.Characters, err error) {
	characters, err = ci.repo.FindAll()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return
}

// Return character by ID
func (ci CharactersInteractor) Show(characterID int) (character *domain.Character, err error) {
	character, err = ci.repo.FindByID(characterID)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return
}
