package user

import (
	"go-clean-arch/entities"
)

type IUsersRepository interface {
	FindAll() ([]entities.User, error)
	Create(user entities.User) error
}
