package usecases

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"go-project/app/domain"
)

type mockCharactersRepo struct {
	mock.Mock
}

func (mr mockCharactersRepo) Index() (characters domain.Characters, err error) {
	arg := mr.Called()

	return arg.Get
}

func TestCharactersInteractor_Index(t *testing.T) {

}
