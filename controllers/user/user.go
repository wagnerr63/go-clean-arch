package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-clean-arch/usecases"
	"go-clean-arch/usecases/user"
)

type IUserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	ListAll(w http.ResponseWriter, r *http.Request)
}

type controllers struct {
	usecases *usecases.Container
}

func New(usecases *usecases.Container) IUserController {
	return &controllers{usecases: usecases}
}

func (ctr *controllers) Create(w http.ResponseWriter, r *http.Request) {
	var userData user.ICreateUserUseCaseDTO
	json.NewDecoder(r.Body).Decode(&userData)

	fmt.Println(userData)

	err := ctr.usecases.User.Create(userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ctr *controllers) ListAll(w http.ResponseWriter, r *http.Request) {
	users, err := ctr.usecases.User.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(users)
}
