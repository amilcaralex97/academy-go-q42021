package repository

import (
	"go-project/app/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

type mockCsv struct {
	mock.Mock
}

var cr = CharactersRepo{
	csvC: mockCsv{},
}

func (mc mockCsv) ReadCsvFiletoString(filePath string) ([][]string, error) {
	return cases, nil
}

func (mc mockCsv) Addrows(characters domain.Characters, filePath string) error {
	return nil
}

func TestCharacterInteractor_FindAll(t *testing.T) {
	testCases := []struct {
		name       string
		response   domain.Characters
		repoErr    error
		expected   domain.Characters
		expectsErr bool
	}{
		{
			"Get charcaters",
			characters,
			nil,
			characters,
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockC := mockCsv{}
			mockC.On("ReadCsvFiletoString").Return(tc.response, tc.repoErr)

			repo := NewCharacterRepo(mockC, "", "")
			actual, err := repo.FindAll()

			if !tc.expectsErr {
				assert.Equal(t, tc.expected, actual)
			}

			if tc.expectsErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
