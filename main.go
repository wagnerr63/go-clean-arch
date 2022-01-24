package main

import (
	"fmt"
	"go-clean-arch/repositories"
	"go-clean-arch/usecases"
	"go-clean-arch/usecases/user"
	"go-clean-arch/configs"
)

func main() {
	repositories := repositories.New()

	usecases := usecases.New(usecases.Options{Repo: repositories})

	usecases.User.Create(user.ICreateUserUseCaseDTO{Name: "Wagner", Email: "wagner@mail.com", Password: "mudar@123"})

	fmt.Println(usecases.User.FindAll())

	router := configs.InitTools()

	router.SERVE(":3333")
}
