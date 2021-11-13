package usecases

import (
	"go-project/app/domain"
)

type CharacterRepository interface {
	FindAll() (domain.Characters, error)
	FindByID(int) (domain.Character, error)
}