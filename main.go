package main

import (
	"go-clean-arch/configs"
	"go-clean-arch/controllers"
	"go-clean-arch/repositories"
	"go-clean-arch/usecases"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repositories := repositories.New(repositories.Options{
		WriterGorm: configs.GetWriterGorm(),
		ReaderGorm: configs.GetReaderGorm(),
	})

	usecases := usecases.New(usecases.Options{Repo: repositories})

	controllers := controllers.New(controllers.Options{Usecases: usecases})

	router := configs.InitTools()

	router.POST("/users", controllers.User.Create)
	router.GET("/users", controllers.User.ListAll)

	router.SERVE(":3333")
}
