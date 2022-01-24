package main

import (
	"fmt"
	"go-clean-arch/repositories"
	"go-clean-arch/usecases"
	"go-clean-arch/usecases/user"
)

func main() {
	repositories := repositories.New()

	usecases := usecases.New(usecases.Options{Repo: repositories})

	usecases.User.Create(user.ICreateUserUseCaseDTO{Name: "Wagner", Email: "wagner@mail.com", Password: "mudar@123"})

	fmt.Println(usecases.User.FindAll())
}
