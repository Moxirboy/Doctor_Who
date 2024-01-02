package rest

import (
	"DoctorWho/internal/domain"
	"DoctorWho/internal/usecase"
	"encoding/json"
	"net/http"
)

type controller struct {
	usecase usecase.Usecase
}

func NewController(usecase usecase.Usecase) *controller {
	return &controller{usecase: usecase}
}

func (c controller) SignUp(w http.ResponseWriter, r *http.Request) {
	var NewUser domain.NewUser
	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		panic(err)
	}
	id, err := c.usecase.RegisterUser(&NewUser)
	//writing to the HEADER
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}
func (c controller) Login(w http.ResponseWriter, r *http.Request) {
	var User domain.User
	err
}
