package repositories

import (
	"go-clean-arch/repositories/user"

	"gorm.io/gorm"
)

type Container struct {
	User user.IUsersRepository
}

type Options struct {
	WriterGorm *gorm.DB
	ReaderGorm *gorm.DB
}

func New(opts Options) *Container {
	return &Container{
		User: user.NewGormRepository(opts.WriterGorm, opts.ReaderGorm),
	}
}
