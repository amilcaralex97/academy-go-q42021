package usecases

import (
	"errors"

	"go-project/app/domain"
	r "go-project/app/repository"
)

type repository interface {
	FindAll(data [][]string) (domain.Characters, error)
	FindByID(data [][]string, characterID int) (*domain.Character, error)
}

type apiI interface {
	FetchCharacters() (domain.Characters, error)
}

type csvI interface {
	ReadCsvFile() ([][]string, error)
}
type CharactersInteractor struct {
	repo repository
	api  apiI
	csv  csvI
}

//NewCharactersInteractor factory character interactor
func NewCharactersInteractor(repo r.CharactersRepo, apiRepo r.ApiRepo, csvRepo r.CsvRepo) CharactersInteractor {
	return CharactersInteractor{repo, apiRepo, csvRepo}
}

//FetchCharacters return fetched characters
func (ci CharactersInteractor) FetchCharacters() (characters domain.Characters, err error) {
	characters, err = ci.api.FetchCharacters()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	data, err := ci.csv.ReadCsvFile()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	lastId := len(data) - 1

	for i := 0; i < len(characters); i++ {
		characters[i].ID = lastId + 1
		lastId++
	}

	return
}

//Index return all characters
func (ci CharactersInteractor) Index() (characters domain.Characters, err error) {
	data, err := ci.csv.ReadCsvFile()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	characters, err = ci.repo.FindAll(data)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return
}

//Show return character by ID
func (ci CharactersInteractor) Show(characterID int) (character *domain.Character, err error) {
	data, err := ci.csv.ReadCsvFile()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	character, err = ci.repo.FindByID(data, characterID)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return
}
