package domain

import (
	"testing"
)

var characters = Characters{
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

func TestModel_CreateCharacter(t *testing.T) {
	for idx, tt := range cases {
		expected := characters[idx]
		character, _ := CreateCharacter(tt)
		if character != expected {
			t.Error("Error while creating character")
		}
	}
}

func TestModel_CreateCharacterList(t *testing.T) {
	charactersTest := CreateCharacterList(cases)

	for idx, tt := range charactersTest {
		expected := characters[idx]

		if tt != expected {
			t.Error("Error while creating character")
		}
	}
}
