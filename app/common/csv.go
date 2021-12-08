package common

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"go-project/app/domain"
)

//const CSVFile string = "/Users/alejandrosanchez/Documents/go_bootcamp/app/resources/characters.csv"

type CsvRepo struct {
}

//NewCharacterRepo factory Character repo
func NewCsvRepo() CsvRepo {
	return CsvRepo{}
}

//ReadCsvFiletoString reads csv File return [][]strings
func (cs CsvRepo) ReadCsvFiletoString(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
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
func (cs CsvRepo) Addrows(characters domain.Characters, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return errors.New(err.Error())
	}
	w := csv.NewWriter(f)
	for _, character := range characters {
		row := []string{strconv.Itoa(character.ID), character.Name}
		w.Write(row)
	}
	w.Flush()

	return nil
}
