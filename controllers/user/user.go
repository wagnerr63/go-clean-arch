package user

import (
	"encoding/json"
	"net/http"

	"go-clean-arch/usecases"
	"go-clean-arch/usecases/user"

	"github.com/gorilla/mux"
)

type IUserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	ListAll(w http.ResponseWriter, r *http.Request)
	GetInfo(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Auth(w http.ResponseWriter, r *http.Request)
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

func (ctr *controllers) GetInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]

	user, err := ctr.usecases.User.GetInfo(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (ctr *controllers) Update(w http.ResponseWriter, r *http.Request) {
	var userData user.IUpdateUserUseCaseDTO
	params := mux.Vars(r)

	userId := params["id"]

	json.NewDecoder(r.Body).Decode(&userData)
	userData.ID = userId

	err := ctr.usecases.User.Update(userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (ctr *controllers) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId := params["id"]
	ctr.usecases.User.Delete(userId)

	w.WriteHeader(http.StatusNoContent)
}

func (ctr *controllers) Auth(w http.ResponseWriter, r *http.Request) {
	var credentials user.IAuthUserUseCaseDTO
	json.NewDecoder(r.Body).Decode(&credentials)

	session, err := ctr.usecases.User.Auth(credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(session)

}
