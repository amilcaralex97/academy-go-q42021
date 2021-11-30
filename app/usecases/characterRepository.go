package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"go-project/app/common"
	"go-project/app/domain"
)

const CSVFile string = "/Users/alejandrosanchez/Documents/go_bootcamp/app/resources/characters.csv"

type CharactersRepo struct{}

func NewCharacterRepo() CharactersRepo {
	return CharactersRepo{}
}

func (CharactersRepo) FetchCharacters() (*domain.Characters, error) {
	resp, err := http.Get("https://swapi.dev/api/people")

	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()

	//Create a variable of the same type as our model
	var cResp *domain.Characters

	json_resp := json.NewDecoder(resp.Body).Decode(&cResp)

	fmt.Println(json_resp)

	//Invoke the text output function & return it with nil as the error value
	return nil, json_resp
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
		}
	}

	if (domain.Character{}) == character {
		return nil, errors.New("error: character doesn't exist")
	}

	return &character, nil
}
