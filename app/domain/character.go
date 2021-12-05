package domain

import (
	"strconv"
)

type Characters []Character

type Character struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Height    int    `json:"height"`
	Mass      int    `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
}

//CreateCharacterList returns a list of characters from CSV
func CreateCharacterList(data [][]string) Characters {
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
				} else if j == 2 {
					var err error
					rec.Height, err = strconv.Atoi(field)
					if err != nil {
						continue
					}
				} else if j == 3 {
					var err error
					rec.Mass, err = strconv.Atoi(field)
					if err != nil {
						continue
					}
				} else if j == 4 {
					rec.HairColor = field
				} else if j == 5 {
					rec.SkinColor = field
				} else if j == 6 {
					rec.EyeColor = field
				} else if j == 7 {
					rec.BirthYear = field
				} else if j == 8 {
					rec.Gender = field
				}
			}
			characterList = append(characterList, rec)
		}
	}
	return characterList
}

//func CharacterCsvWriter() {
//    writer := csv.Writer()
//}
