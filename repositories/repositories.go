package repositories

import "go-clean-arch/repositories/user"

type Container struct {
	User user.IUsersRepository
}

func New() *Container {
	return &Container{
		User: user.NewMockUsersRepository(),
	}
}
