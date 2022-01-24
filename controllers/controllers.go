package controllers

import (
	"go-clean-arch/controllers/user"
	"go-clean-arch/usecases"
)

type Container struct {
	User user.IUserController
}

type Options struct {
	Usecases *usecases.Container
}

func New(opts Options) *Container {
	return &Container{
		User: user.New(opts.Usecases),
	}
}
