package domain

import (
	"strconv"
)

type Characters []Character

type Character struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func CreateCharacterList(data [][]string) []Character {
    var characterList []Character
    for i, line := range data {
        if i > 0 {
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