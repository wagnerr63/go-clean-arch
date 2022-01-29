package usecases

import (
	"go-clean-arch/repositories"
	"go-clean-arch/usecases/user"
	"go-clean-arch/utils/token"
)

type Container struct {
	User user.IUsersUseCases
}

type Options struct {
	Repo  *repositories.Container
	Token token.TokenHash
}

func New(opts Options) *Container {

	return &Container{
		User: user.New(opts.Repo, opts.Token),
	}
}
