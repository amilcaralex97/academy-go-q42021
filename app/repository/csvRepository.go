package repository

import (
	"encoding/csv"
	"errors"
	"os"
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
