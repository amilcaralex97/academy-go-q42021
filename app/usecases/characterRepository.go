package usecases

import (
	"fmt"

	"go-project/app/common"
	"go-project/app/domain"
)

func FindAll() (domain.Characters) {
	data := common.ReadCsvFile("/home/amilcar/Documents/Projects/go-project/app/resources/characters.csv")

	characterList := common.CreateCharacterList(data)

	fmt.Println(characterList)

	return characterList
}

func FindByID(characterID int) (domain.Characters) {
	data := common.ReadCsvFile("/home/amilcar/Documents/Projects/go-project/app/resources/characters.csv")

	characterList := common.CreateCharacterList(data)

	for k, v := range characterList {
		fmt.Println(k, ' ', v)
	}

	fmt.Println(characterList, characterID)

	return characterList
}