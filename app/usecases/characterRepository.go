package usecases

import (
	"errors"

	"go-project/app/common"
	"go-project/app/domain"
)

const CSVFile string = "/home/amilcar/Documents/Projects/go-project/app/resources/characters.csv"
const CsvError string = "error while trying to read CSV file"

type CharactersRepo struct{}


func NewCharacterRepo() CharactersRepo{
	return CharactersRepo{}
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
		if(v.ID == characterID){
			character = v
		}
	}

	if (domain.Character{}) == character {
		return nil, errors.New("error: character doesn't exist")
	}

	return &character, nil
}