package repository

import (
	"reflect"
	"testing"
)

func TestModel_WorkerPoolCsv(t *testing.T) {
	charactersTest, _ := NewWorkerPool().WorkerPoolCsv("odd", 6, 5)
	expectedLen := 6

	if reflect.TypeOf(charactersTest[0]) != reflect.TypeOf(characters[0]) {
		t.Error("response is not type Character")
	}

	if len(charactersTest) != expectedLen {
		t.Error("response wrong size")
	}
}
