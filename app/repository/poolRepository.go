package repository

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"sync"

	"go-project/app/domain"
)

const CSVFile string = "/Users/alejandrosanchez/Documents/go_bootcamp/app/resources/characters.csv"

type WorkerPool struct {
}

// NewWorkerPool will create an instance of WorkerPool.
func NewWorkerPool() WorkerPool {
	return WorkerPool{}
}

func (wp WorkerPool) worker(t string, jobs <-chan []string, results chan<- domain.Character) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			character := domain.CreateCharacter(job)
			if t == "odd" && character.ID%2 != 0 {
				results <- character
			} else if t == "even" && character.ID%2 == 0 {
				results <- character
			}
		}
	}
}

func (wp WorkerPool) WorkerPoolCsv(t string, items int, itpw int) (domain.Characters, error) {
	file, err := os.Open(CSVFile)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer file.Close()

	reader := csv.NewReader(file)

	var characters domain.Characters

	workers := items / itpw

	jobs := make(chan []string, items)
	res := make(chan domain.Character, items)

	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			wp.worker(t, jobs, res)
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

	return characters, nil
}
