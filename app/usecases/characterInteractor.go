package usecases

import (
	"errors"

	"go-project/app/domain"
)

type repository interface {
	FindAll() (domain.Characters, error)
	FetchCharacters() (domain.Characters, error)
	WorkerPoolCsv(t string, items int, itpw int) (domain.Characters, error)
}

type CharactersInteractor struct {
	repo repository
}

//NewCharactersInteractor factory character interactor
func NewCharactersInteractor(repo repository) CharactersInteractor {
	return CharactersInteractor{repo}
}

//Return characters concurrently
func (ci CharactersInteractor) CharactersConcurrently(t string, items int, itpw int) (characters domain.Characters, err error) {
	characters, err = ci.repo.WorkerPoolCsv(t, items, itpw)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return characters, nil
}

//FetchCharacters return fetched characters
func (ci CharactersInteractor) FetchCharacters() (characters domain.Characters, err error) {
	characters, err = ci.repo.FetchCharacters()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return
}

//Index return all characters
func (ci CharactersInteractor) Index() (characters domain.Characters, err error) {
	characters, err = ci.repo.FindAll()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return
}

//Show return character by ID
func (ci CharactersInteractor) Show(characterID int) (character domain.Character, err error) {
	characters, err := ci.repo.FindAll()

	if err != nil {
		return domain.Character{}, errors.New(err.Error())
	}

	character = domain.Character{}

	for _, v := range characters {
		if v.ID == characterID {
			character = v
			break
		}
	}

	if (domain.Character{}) == character {
		return domain.Character{}, errors.New("error: character doesn't exist")
	}

	return
}
