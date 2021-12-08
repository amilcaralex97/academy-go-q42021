package usecases

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

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

var charactersConc = domain.Characters{
	{
		ID:   3,
		Name: "C-3PO",
	},
	{
		ID:   1,
		Name: "Amilcar Sanchez",
	},
	{
		ID:   5,
		Name: "Darth Vader",
	},
	{
		ID:   9,
		Name: "R5-D4",
	},
	{
		ID:   11,
		Name: "Obi-Wan Kenobi",
	},
	{
		ID:   13,
		Name: "C-3PO",
	},
}

type mockCharacterRepository struct {
	mock.Mock
}

func (m mockCharacterRepository) FindAll() (domain.Characters, error) {
	args := m.Called()
	return args.Get(0).(domain.Characters), args.Error(1)
}

func (m mockCharacterRepository) FetchCharacters() (domain.Characters, error) {
	args := m.Called()
	return args.Get(0).(domain.Characters), args.Error(1)
}

func (m mockCharacterRepository) WorkerPoolCsv(t string, items int, itpw int) (domain.Characters, error) {
	args := m.Called(t, items, itpw)
	return args.Get(0).(domain.Characters), args.Error(1)
}

func TestCharacterInteractor_Index(t *testing.T) {
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
		{
			"First character is Amilcar",
			characters,
			nil,
			characters,
			false,
		},
		{
			"error ocurred",
			nil,
			errors.New(""),
			nil,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCharacters := mockCharacterRepository{}
			mockCharacters.On("FindAll").Return(tc.response, tc.repoErr)

			service := NewCharactersInteractor(mockCharacters)
			actual, err := service.Index()

			if !tc.expectsErr {
				assert.Equal(t, tc.expected, actual)
				assert.Equal(t, actual[0], tc.expected[0])
			}

			if tc.expectsErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})

	}
}

func TestCharacterInteractor_Show(t *testing.T) {
	testCases := []struct {
		name       string
		response   domain.Characters
		id         int
		repoErr    error
		expected   domain.Character
		expectsErr bool
	}{
		{
			"Get charcater",
			characters,
			1,
			nil,
			characters[0],
			false,
		},
		{
			"Character Not found",
			characters,
			4,
			errors.New(""),
			domain.Character{},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCharacters := mockCharacterRepository{}
			mockCharacters.On("FindAll").Return(tc.response, tc.repoErr)

			service := NewCharactersInteractor(mockCharacters)
			actual, err := service.Show(tc.id)

			fmt.Println(actual)

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

func TestCharacterInteractor_FetchCharacters(t *testing.T) {
	testCases := []struct {
		name       string
		response   domain.Characters
		repoErr    error
		expected   domain.Characters
		expectsErr bool
	}{
		{
			"Get fetched charcaters",
			characters,
			nil,
			characters,
			false,
		},
		{
			"First character is Amilcar",
			characters,
			nil,
			characters,
			false,
		},
		{
			"error ocurred",
			nil,
			errors.New(""),
			nil,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCharacters := mockCharacterRepository{}
			mockCharacters.On("FetchCharacters").Return(tc.response, tc.repoErr)

			service := NewCharactersInteractor(mockCharacters)
			actual, err := service.FetchCharacters()

			if !tc.expectsErr {
				assert.Equal(t, tc.expected, actual)
				assert.Equal(t, actual[0], tc.expected[0])
			}

			if tc.expectsErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestCharacterInteractor_CharactersConcurrently(t *testing.T) {
	testCases := []struct {
		name       string
		response   domain.Characters
		repoErr    error
		expected   domain.Characters
		expectsErr bool
	}{
		{
			"Get characters concurrently",
			charactersConc,
			nil,
			charactersConc,
			false,
		},
		{
			"error ocurred",
			nil,
			errors.New(""),
			nil,
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCharacters := mockCharacterRepository{}
			mockCharacters.On("WorkerPoolCsv", "odd", 6, 5).Return(tc.response, tc.repoErr)

			service := NewCharactersInteractor(mockCharacters)
			actual, err := service.CharactersConcurrently("odd", 6, 5)

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
