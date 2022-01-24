package user

import (
	"go-clean-arch/entities"
	"go-clean-arch/repositories"
)

type ICreateUserUseCaseDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type IUsersUseCases interface {
	Create(data ICreateUserUseCaseDTO) error
	FindAll() ([]entities.User, error)
}

type userUserCases struct {
	repositories *repositories.Container
}

func New(repo *repositories.Container) IUsersUseCases {
	return &userUserCases{repositories: repo}
}

func (usecase *userUserCases) Create(data ICreateUserUseCaseDTO) error {
	user := entities.NewUser()
	user.Name = data.Name
	user.Email = data.Email

	// TODO: hash password
	user.Password = data.Password

	return usecase.repositories.User.Create(user)
}

func (usecase *userUserCases) FindAll() ([]entities.User, error) {
	return usecase.repositories.User.FindAll()
}
