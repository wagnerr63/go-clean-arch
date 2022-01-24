package usecases

import (
	"go-clean-arch/repositories"
	"go-clean-arch/usecases/user"
)

type Container struct {
	User user.IUsersUseCases
}

type Options struct {
	Repo *repositories.Container
}

func New(opts Options) *Container {

	return &Container{
		User: user.New(opts.Repo),
	}
}
