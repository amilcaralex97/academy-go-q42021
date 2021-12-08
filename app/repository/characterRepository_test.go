package repository

import (
	"reflect"
	"testing"

	"go-project/app/domain"
)

var characters = domain.Characters{
	{
		ID:   1,
		Name: "Amilcar Sanchez",
	},
	{
		ID:   2,
		Name: "Luke Skywalker",
	},
}

var cases = [][]string{
	{
		"1", "Amilcar Sanchez",
	},
	{
		"2", "Luke Skywalker",
	},
}

func TestModel_FindAll(t *testing.T) {
	charactersTest, _ := NewCharacterRepo().FindAll(cases)
	expectedLen := 2

	if len(charactersTest) != expectedLen {
		t.Error("characters wrong size")
	}
}

func TestModel_WorkerPoolCsv(t *testing.T) {
	charactersTest, _ := NewCharacterRepo().WorkerPoolCsv("odd", 6, 5)
	expectedLen := 6

	if reflect.TypeOf(charactersTest[0]) != reflect.TypeOf(characters[0]) {
		t.Error("response is not type Character")
	}

	if len(charactersTest) != expectedLen {
		t.Error("response wrong size")
	}
}

func TestModel_FetchCharacters(t *testing.T) {
	charactersTest, _ := NewCharacterRepo().FetchCharacters()
	expectedLen := 10

	if reflect.TypeOf(charactersTest[0]) != reflect.TypeOf(characters[0]) {
		t.Error("response is not type Character")
	}

	if len(charactersTest) != expectedLen {
		t.Error("response wrong size")
	}
}
