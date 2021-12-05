package repository

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"go-project/app/domain"
)

//const CSVFile string = "/Users/alejandrosanchez/Documents/go_bootcamp/app/resources/characters.csv"

type CsvRepo struct {
	filePath string
}

//NewCharacterRepo factory Character repo
func NewCsvRepo(filePath string) CsvRepo {
	return CsvRepo{filePath: filePath}
}

//ReadCsvFile reads csv File
func (cs CsvRepo) ReadCsvFile() ([][]string, error) {
	f, err := os.Open(cs.filePath)

	if err != nil {
		return nil, errors.New("error while trying to open CSV file")
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return records, nil
}

// Updates Csvfile
func (cs CsvRepo) Addrow(characters domain.Characters) error {
	f, err := os.OpenFile(cs.filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return errors.New(err.Error())
	}
	w := csv.NewWriter(f)
	for _, character := range characters {
		row := []string{strconv.Itoa(character.ID), character.Name, strconv.Itoa(character.Height), strconv.Itoa(character.Mass), character.HairColor, character.SkinColor, character.EyeColor, character.BirthYear, character.Gender}
		w.Write(row)
	}
	w.Flush()

	return nil
}
