package repository

import (
	"encoding/csv"
	"errors"
	"io"
	"math"
	"os"
	"strconv"
	"sync"

	"go-project/app/domain"
)

const CSVFile string = "/Users/alejandrosanchez/Documents/go_bootcamp/app/resources/characters.csv"

var counter = 0

type WorkerPool struct {
}

// NewWorkerPool will create an instance of WorkerPool.
func NewWorkerPool() WorkerPool {
	return WorkerPool{}
}

func (wp WorkerPool) worker(t string, jobs <-chan []string, results chan domain.Character, items int) {
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

func (wp WorkerPool) WorkerPoolCsv(t string, items int, itpw int) (domain.Characters, error) {
	file, err := os.Open(CSVFile)

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
			wp.worker(t, jobs, res, items)
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
