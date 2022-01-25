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
