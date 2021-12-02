package usecases

import (
	"errors"

	"go-project/app/domain"
)

type repository interface {
	FindAll() (*domain.Characters, error)
	FindByID(characterID int) (*domain.Character, error)
	FetchCharacters() (*domain.Characters, error)
}

type CharactersInteractor struct {
	repo repository
}

//NewCharactersInteractor factory character interactor
func NewCharactersInteractor(repo CharactersRepo) CharactersInteractor {
	return CharactersInteractor{repo}
}

//FetchCharacters return fetched characters
func (ci CharactersInteractor) FetchCharacters() (characters *domain.Characters, err error) {
	characters, err = ci.repo.FetchCharacters()

	return
}

//Index return all characters
func (ci CharactersInteractor) Index() (characters *domain.Characters, err error) {
	characters, err = ci.repo.FindAll()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return
}

//Show return character by ID
func (ci CharactersInteractor) Show(characterID int) (character *domain.Character, err error) {
	character, err = ci.repo.FindByID(characterID)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return
}
