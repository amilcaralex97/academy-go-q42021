package usecases

import (
	"errors"
	"fmt"

	"go-project/app/common"
	"go-project/app/domain"
)

const CSVFile string = "/home/amilcar/Documents/Projects/go-project/app/resources/characters.csv"
const CsvError string = "error while trying to read CSV file"

// type characterRepository interface {
// 	FindAll() (domain.Characters, error)
// 	FindByID(int) (domain.Character, error)
// }

func FindAll() (*domain.Characters, error) {
	data, err := common.ReadCsvFile(CSVFile)

	if err != nil {
		return nil, errors.New(CsvError)
	}

	characterList := domain.CreateCharacterList(data)

	return &characterList, nil
}

func FindByID(characterID int) (*domain.Character, error) {
	data, err := common.ReadCsvFile(CSVFile)

	fmt.Println(characterID)

	if err != nil {
		return nil, errors.New(CsvError)
	}

	characterList := domain.CreateCharacterList(data)

	fmt.Println(characterList)

	var character domain.Character

	for _, v := range characterList {
		if(v.ID == characterID){
			character = v
		}
	} 

	fmt.Println(character)

	return &character, nil
}