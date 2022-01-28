package user

import (
	"errors"
	"go-clean-arch/entities"

	"gorm.io/gorm"
)

type repoGorm struct {
	writer *gorm.DB
	reader *gorm.DB
}

func NewGormRepository(writer, reader *gorm.DB) IUsersRepository {
	return &repoGorm{writer: writer, reader: reader}
}

func (repo *repoGorm) Create(user entities.User) error {

	repo.writer.Table("users").Create(&user)

	return repo.writer.Error
}

func (repo *repoGorm) FindAll() ([]entities.User, error) {
	var users []entities.User
	repo.reader.Table("users").Find(&users)

	if repo.reader.Error != nil {
		return nil, errors.New("Users not found")
	}

	return users, nil
}

func (repo *repoGorm) FindById(id string) (entities.User, error) {
	var userById entities.User

	repo.reader.Table("users").Find(&userById)

	if repo.reader.Error != nil {
		return entities.User{}, errors.New("User not found.")
	}

	return userById, nil
}

func (repo *repoGorm) FindByEmail(email string) (entities.User, error) {
	var userByEmail entities.User

	repo.reader.Table("users").Where("email = ?", email).Find(&userByEmail)

	if repo.reader.Error != nil {
		return entities.User{}, errors.New("User not found.")
	}

	return userByEmail, nil
}

func (repo *repoGorm) Update(user entities.User) error {
	repo.writer.Table("users").Updates(user)

	if repo.writer.Error != nil {
		return errors.New("User update error")
	}

	return nil
}

func (repo *repoGorm) Delete(id string) error {
	repo.writer.Table("users").Delete(id)

	if repo.writer.Error != nil {
		return errors.New("User delete error")
	}

	return nil
}
