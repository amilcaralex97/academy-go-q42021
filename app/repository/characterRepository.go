package repository

import (
	"errors"

	"go-project/app/domain"
)

type CharactersRepo struct{}

type Response struct {
	Results domain.Characters `json:"results"`
}

//NewCharacterRepo factory Character repo
func NewCharacterRepo() CharactersRepo {
	return CharactersRepo{}
}

//FindAll gets characters from csv
func (CharactersRepo) FindAll(data [][]string) (domain.Characters, error) {

	characterList := domain.CreateCharacterList(data)

	return characterList, nil
}

//FindByID get character in the csv by ID
func (CharactersRepo) FindByID(data [][]string, characterID int) (*domain.Character, error) {
	characterList := domain.CreateCharacterList(data)

	character := domain.Character{}

	for _, v := range characterList {
		if v.ID == characterID {
			character = v
			break
		}
	}

	if (domain.Character{}) == character {
		return nil, errors.New("error: character doesn't exist")
	}

	return &character, nil
}
