package usecases

import (
	"errors"
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
	{
		ID:   3,
		Name: "C-3PO",
	},
	{
		ID:   4,
		Name: "R2-D2",
	},
}

type mockCharactersRepo struct {
	mock.Mock
}

type mockApiRepo struct {
	mock.Mock
}

type mockCsvRepo struct {
	mock.Mock
}

type mockPoolRepo struct {
	mock.Mock
}

func (mr mockCharactersRepo) Index() (domain.Characters, error) {
	arg := mr.Called()

	return arg.Get(0).(domain.Characters), arg.Error(1)
}

func (mp mockPoolRepo) CharactersConcurrently(domain.Characters, error) {
	arg := mp.Called()

	return arg.Get(0).(domain.Characters), arg.Error(1)
}

func TestCharactersInteractor_CharactersConcurrently(t *testing.T) {
	testCases := []struct {
		name     string
		response domain.Characters
		expected domain.Characters
		err      error
	}{
		{
			"get all characters",
			characters,
			characters,
			nil,
		},
		{
			"error in repository",
			domain.Characters{},
			domain.Characters{},
			errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mockCharactersRepo{}
			mockApi := mockApiRepo{}
			mockCsv := mockCsvRepo{}
			mockPool := mockCsvRepo{}
			mock.On("FindAll").Return(tc.response, tc.err)

			service := NewCharactersInteractor(mock, mockApi, mockCsv, mockPool)

			index, err := service.Index()

			assert.Equal(t, tc.expected, index)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestCharactersInteractor_Index(t *testing.T) {
	testCases := []struct {
		name     string
		response domain.Characters
		expected domain.Characters
		err      error
	}{
		{
			"get all characters",
			characters,
			characters,
			nil,
		},
		{
			"error in repository",
			domain.Characters{},
			domain.Characters{},
			errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mockCharactersRepo{}
			mockApi := mockApiRepo{}
			mockCsv := mockCsvRepo{}
			mockPool := mockCsvRepo{}
			mock.On("FindAll").Return(tc.response, tc.err)

			service := NewCharactersInteractor(mock, mockApi, mockCsv, mockPool)

			index, err := service.Index()

			assert.Equal(t, tc.expected, index)
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestCharactersInteractor_Show(t *testing.T) {
	expected := domain.Character{
		1,
		"Amilcar Sanchez",
	}

	mockRepo := mockCharactersRepo{}
	mockApi := mockApiRepo{}
	mockCsv := mockCsvRepo{}
	mockPool := mockCsvRepo{}
	mockRepo.On("FindByID").Return(characters, nil)

	service := NewCharactersInteractor(mockRepo, mockApi, mockCsv, mockPool)
	res, err := service.Show(1)

	assert.Equal(t, expected, res)
	assert.Equal(t, err, nil)
}

func TestCharactersInteractor_FetchCharacters(t *testing.T) {
	testCases := []struct {
		name     string
		response domain.Characters
		expected domain.Characters
		err      error
	}{
		{
			"get all characters",
			characters,
			characters,
			nil,
		},
		{
			"error in repository",
			domain.Characters{},
			domain.Characters{},
			errors.New("test error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mockCharactersRepo{}
			mockApi := mockApiRepo{}
			mockCsv := mockCsvRepo{}
			mockPool := mockCsvRepo{}
			mock.On("FetchCharacters").Return(tc.response, tc.err)

			service := NewCharactersInteractor(mock, mockApi, mockCsv, mockPool)

			index, err := service.Index()

			assert.Equal(t, tc.expected, index)
			assert.Equal(t, tc.err, err)
		})
	}
}
