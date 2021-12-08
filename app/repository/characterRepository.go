package repository

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"sync"

	"go-project/app/domain"
)

var counter = 0

type csvI interface {
	ReadCsvFiletoString(filePath string) ([][]string, error)
	Addrows(characters domain.Characters, filePath string) error
}

type Response struct {
	Results domain.Characters `json:"results"`
}

type CharactersRepo struct {
	csvC csvI
	file string
	api  string
}

//NewCharacterRepo factory Character repo
func NewCharacterRepo(csvC csvI, file string, api string) CharactersRepo {
	return CharactersRepo{csvC, file, api}
}

//FindAll gets characters from csv
func (cr CharactersRepo) FindAll() (domain.Characters, error) {
	data, err := cr.csvC.ReadCsvFiletoString(cr.file)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	characterList := domain.CreateCharacterList(data)

	return characterList, nil
}

//FetchCharacters gets characters from an API
func (cr CharactersRepo) FetchCharacters() (domain.Characters, error) {
	response, err := http.Get(cr.api)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	responseObject := Response{}
	json.Unmarshal(responseData, &responseObject)

	data, err := cr.csvC.ReadCsvFiletoString(cr.file)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	lastId := len(data)

	for i := 0; i < len(responseObject.Results); i++ {
		responseObject.Results[i].ID = lastId + 1
		lastId++
	}

	cr.csvC.Addrows(responseObject.Results, cr.file)

	charactersCsv, err := cr.FindAll()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	charactersCsv = append(charactersCsv, responseObject.Results...)

	return charactersCsv, nil
}

// WorkerPoolCsv csv reader worker pool
func (cr CharactersRepo) WorkerPoolCsv(t string, items int, itpw int) (domain.Characters, error) {
	file, err := os.Open(cr.file)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer file.Close()

	reader := csv.NewReader(file)

	var characters domain.Characters

	workers := int(math.Ceil(float64(items) / float64(itpw)))

	jobs := make(chan []string)
	res := make(chan domain.Character, items-1)

	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			cr.worker(t, jobs, res, items)
		}()
	}

	go func() {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				continue
			}
			jobs <- record
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		characters = append(characters, r)
	}

	counter = 0

	return characters, nil
}

// worker worker from WorkerPoolCsv
func (CharactersRepo) worker(t string, jobs <-chan []string, results chan domain.Character, items int) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}

			if counter == items {
				return
			}

			id, _ := strconv.Atoi(job[0])

			if t == "odd" && id%2 == 0 {
				continue
			} else if t == "even" && id%2 != 0 {
				continue
			}

			character, err := domain.CreateCharacter(job)

			if err != nil {
				continue
			}

			results <- character
		}
		counter++
	}
}
