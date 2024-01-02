package main

import (
	configs "DoctorWho/internal/common/config"
	"DoctorWho/internal/delivery/rest"
	"DoctorWho/internal/repository"
	"DoctorWho/internal/usecase"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	instance := configs.Configuration()
	postgres, err := configs.NewPostgresConfig(*instance)
	if err != nil {
		panic("no connection")
	}
	repo := repository.NewRepo(postgres)
	service := usecase.NewUserUsecase(repo)
	controller := rest.NewController(service)
	router := mux.NewRouter()
	router.HandleFunc("/signUp", controller.SignUp)
	server := &http.Server{Addr: ":5005", Handler: router}
	server.ListenAndServe()
}
