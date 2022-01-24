package main

import (
	"go-clean-arch/configs"
	"go-clean-arch/controllers"
	"go-clean-arch/repositories"
	"go-clean-arch/usecases"
)

func main() {
	repositories := repositories.New()

	usecases := usecases.New(usecases.Options{Repo: repositories})

	controllers := controllers.New(controllers.Options{Usecases: usecases})

	router := configs.InitTools()

	router.POST("/users", controllers.User.Create)
	router.GET("/users", controllers.User.ListAll)

	router.SERVE(":3333")
}
