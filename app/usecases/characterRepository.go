package usecases

import (
	"fmt"

	"go-project/app/common"
	"go-project/app/domain"
)

const CSVFile string = "/home/amilcar/Documents/Projects/go-project/app/resources/characters.csv"

func FindAll() (domain.Characters, error) {
	data := common.ReadCsvFile(CSVFile)

	characterList := domain.CreateCharacterList(data)

	return characterList, nil
}

func FindByID(characterID int) (domain.Character, error) {
	data := common.ReadCsvFile(CSVFile)

	characterList := domain.CreateCharacterList(data)

	fmt.Println(characterList)

	var character domain.Character

	for _, v := range characterList {
		if(v.ID == characterID){
			character = v
		}
	}

	return character, nil
}