package repository

import (
	"reflect"
	"testing"
)

func TestModel_FetchCharacters(t *testing.T) {
	charactersTest, _ := NewApiRepo("https://swapi.dev/api/people").FetchCharacters()
	expectedLen := 10

	if reflect.TypeOf(charactersTest[0]) != reflect.TypeOf(characters[0]) {
		t.Error("response is not type Character")
	}

	if len(charactersTest) != expectedLen {
		t.Error("response wrong size")
	}
}
