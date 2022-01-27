package user

import "go-clean-arch/entities"

type MockUsersRepository struct {
	users []entities.User
}

func NewMockUsersRepository() IUsersRepository {
	return &MockUsersRepository{users: nil}
}

func (repo *MockUsersRepository) FindAll() ([]entities.User, error) {
	return repo.users, nil
}

func (repo *MockUsersRepository) Create(user entities.User) error {
	repo.users = append(repo.users, user)
	return nil
}

func (repo *MockUsersRepository) FindById(id string) (entities.User, error) {
	return entities.User{}, nil
}

func (repo *MockUsersRepository) FindByEmail(id string) (entities.User, error) {
	return entities.User{}, nil
}

func (repo *MockUsersRepository) Update(user entities.User) error {
	return nil
}
