package user

import (
	"errors"
	"go-clean-arch/entities"
	"go-clean-arch/repositories"
	"go-clean-arch/utils/hasher"
	"go-clean-arch/utils/token"
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

type IAuthUserUseCaseDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type IAuthResponse struct {
	User  entities.User
	Token string `json:"token"`
}

type IUsersUseCases interface {
	Create(data ICreateUserUseCaseDTO) error
	FindAll() ([]entities.User, error)
	GetInfo(id string) (entities.User, error)
	Update(data IUpdateUserUseCaseDTO) error
	Delete(id string) error
	Auth(credentials IAuthUserUseCaseDTO) (IAuthResponse, error)
}

type userUserCases struct {
	repositories *repositories.Container
	tokenHash    token.TokenHash
}

func New(repo *repositories.Container, token token.TokenHash) IUsersUseCases {
	return &userUserCases{repositories: repo, tokenHash: token}
}

func (usecase *userUserCases) Create(data ICreateUserUseCaseDTO) error {
	user := entities.NewUser()
	user.Name = data.Name
	user.Email = data.Email

	hasherBcrypt := hasher.NewBcryptHasher()
	passwordHashed, errHash := hasherBcrypt.Generate(data.Password)
	if errHash != nil {
		return errHash
	}
	user.Password = passwordHashed

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

func (usecase *userUserCases) Auth(credentials IAuthUserUseCaseDTO) (IAuthResponse, error) {
	userByEmail, err := usecase.repositories.User.FindByEmail(credentials.Email)
	if err != nil {
		return IAuthResponse{}, errors.New("user not found")
	}

	hasherBcrypt := hasher.NewBcryptHasher()
	err = hasherBcrypt.Compare(userByEmail.Password, credentials.Password)
	if err != nil {
		return IAuthResponse{}, err
	}

	token, err := usecase.tokenHash.Encrypt(userByEmail)
	if err != nil {
		return IAuthResponse{}, err
	}

	return IAuthResponse{User: userByEmail, Token: token}, nil

}
