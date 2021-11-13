package common

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"go-project/app/domain"
)

func ReadCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func CreateCharacterList(data [][]string) []domain.Character {
    var characterList []domain.Character
    for i, line := range data {
        if i > 0 {
            var rec domain.Character
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