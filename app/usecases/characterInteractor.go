package usecases

import (
	"encoding/csv"
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
	ReadCsvFiletoString() ([][]string, error)
	ReadCsvFile() (*csv.Reader, error)
	Addrows(characters domain.Characters) error
}

type workerPool interface {
	WorkerPoolCsv(t string, items int, itpw int) (domain.Characters, error)
}

type CharactersInteractor struct {
	repo repository
	api  apiI
	csv  csvI
	pool workerPool
}

//NewCharactersInteractor factory character interactor
func NewCharactersInteractor(repo r.CharactersRepo, apiRepo r.ApiRepo, csvRepo r.CsvRepo, poolRepo r.WorkerPool) CharactersInteractor {
	return CharactersInteractor{repo, apiRepo, csvRepo, poolRepo}
}

//Return characters concurrently
func (ci CharactersInteractor) CharactersConcurrently(t string, items int, itpw int) (characters domain.Characters, err error) {
	characters, err = ci.pool.WorkerPoolCsv(t, items, itpw)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return characters, nil
}

//FetchCharacters return fetched characters
func (ci CharactersInteractor) FetchCharacters() (charactersCsv domain.Characters, err error) {
	characters, err := ci.api.FetchCharacters()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	data, err := ci.csv.ReadCsvFiletoString()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	lastId := len(data)

	for i := 0; i < len(characters); i++ {
		characters[i].ID = lastId + 1
		lastId++
	}

	ci.csv.Addrows(characters)

	charactersCsv, err = ci.repo.FindAll(data)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	charactersCsv = append(charactersCsv, characters...)

	return
}

//Index return all characters
func (ci CharactersInteractor) Index() (characters domain.Characters, err error) {
	data, err := ci.csv.ReadCsvFiletoString()
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
	data, err := ci.csv.ReadCsvFiletoString()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	character, err = ci.repo.FindByID(data, characterID)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return
}
