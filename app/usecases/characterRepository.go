package usecases

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"go-project/app/common"
	"go-project/app/domain"
)

const CSVFile string = "/Users/alejandrosanchez/Documents/go_bootcamp/app/resources/characters.csv"

type CharactersRepo struct{}

type Response struct {
	Results domain.Characters `json:"results"`
}

func NewCharacterRepo() CharactersRepo {
	return CharactersRepo{}
}

func (CharactersRepo) FetchCharacters() (*domain.Characters, error) {
	response, err := http.Get("https://swapi.dev/api/people")
	if err != nil {
		return nil, errors.New(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	responseObject := Response{}
	json.Unmarshal(responseData, &responseObject)

	data, err := common.ReadCsvFile(CSVFile)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	lastId := len(data) - 1

	for i := 0; i < len(responseObject.Results); i++ {
		responseObject.Results[i].ID = lastId + 1
		lastId++
	}

	return &responseObject.Results, nil
}

func (CharactersRepo) FindAll() (*domain.Characters, error) {
	data, err := common.ReadCsvFile(CSVFile)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	characterList := domain.CreateCharacterList(data)

	return &characterList, nil
}

func (CharactersRepo) FindByID(characterID int) (*domain.Character, error) {
	data, err := common.ReadCsvFile(CSVFile)

	if err != nil {
		return nil, errors.New(err.Error())
	}

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
