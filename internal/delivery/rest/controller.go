package rest

import (
	"DoctorWho/internal/delivery/dto"
	"DoctorWho/internal/domain"
	"DoctorWho/internal/usecase"
	"encoding/json"
	"log"
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
	var User dto.User
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		panic(err)
	}
	exist, err := c.usecase.Login(User.Phone_number)
	if err != nil {
		panic(err)
	}
	log.Println(exist)
	if exist {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("success")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAlreadyReported)
		json.NewEncoder(w).Encode("fuck you")
	}

}
func (c controller) GetAll(w http.ResponseWriter, r *http.Request) {
	user := c.usecase.GetAll()
	log.Println(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}
