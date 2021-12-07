package repository

import (
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

func TestModel_FindByID(t *testing.T) {
	expected := characters[0]
	character, _ := NewCharacterRepo().FindByID(cases, 1)
	if character != expected {
		t.Fail()
		t.Log("Actual:", character, "Expected:", expected)
	}
}
