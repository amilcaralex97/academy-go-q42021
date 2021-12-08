package domain

import (
	"errors"
	"strconv"
)

type Characters []Character

type Character struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//CreateCharacterList returns a list of characters from CSV
func CreateCharacterList(data [][]string) Characters {
	var characterList []Character
	for i, line := range data {
		if i >= 0 {
			var rec Character
			for j, field := range line {
				if j == 0 {
					var err error
					rec.ID, err = strconv.Atoi(field)
					if err != nil {
						continue
					}
				} else if j == 1 {
					rec.Name = field
				}
			}
			characterList = append(characterList, rec)
		}
	}
	return characterList
}

//Parse character
func CreateCharacter(data []string) (character Character, err error) {
	id, err := strconv.Atoi(data[0])

	if err != nil {
		return Character{}, errors.New("error while parsig character id")
	}

	character = Character{
		ID:   id,
		Name: data[1],
	}

	return character, nil
}
