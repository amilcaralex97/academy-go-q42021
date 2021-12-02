package common

import (
	"encoding/csv"
	"errors"
	"os"
)

//ReadCsvFile reads csv File
func ReadCsvFile(filePath string) ([][]string, error) {
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