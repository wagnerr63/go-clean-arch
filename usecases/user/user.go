package user

import (
	"errors"
	"go-clean-arch/entities"
	"go-clean-arch/repositories"
)

type ICreateUserUseCaseDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type IUpdateUserUseCaseDTO struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	IsAdmin  bool   `json:"is_admin,omitempty"`
}

type IUsersUseCases interface {
	Create(data ICreateUserUseCaseDTO) error
	FindAll() ([]entities.User, error)
	GetInfo(id string) (entities.User, error)
	Update(data IUpdateUserUseCaseDTO) error
	Delete(id string) error
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

func (usecase *userUserCases) GetInfo(id string) (entities.User, error) {
	return usecase.repositories.User.FindById(id)
}

func (usecase *userUserCases) Update(data IUpdateUserUseCaseDTO) error {
	_, err := usecase.repositories.User.FindById(data.ID)
	if err != nil {
		return errors.New("user not found")
	}

	err = usecase.repositories.User.Update(entities.User{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		IsAdmin:  data.IsAdmin,
	})

	if err != nil {
		return errors.New("error trying to update user")
	}

	return nil
}

func (usecase *userUserCases) Delete(id string) error {
	_, err := usecase.repositories.User.FindById(id)
	if err != nil {
		return errors.New("user not found")
	}

	return usecase.repositories.User.Delete(id)
}
