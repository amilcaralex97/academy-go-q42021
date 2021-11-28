package common

import (
	"encoding/csv"
	"errors"
	"os"
)

func ReadCsvFile(filePath string) ([][]string, error) {
    f, err := os.Open(filePath)

    if err != nil {
        return nil, errors.New("error while trying to open CSV file")
    }
    
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()

    if err != nil {
        return nil, errors.New("unable to parse file as CSV for")
    }

    return records, nil
}