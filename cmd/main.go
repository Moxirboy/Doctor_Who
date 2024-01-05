package main

import (
	configs "DoctorWho/internal/common/config"
	"DoctorWho/internal/delivery/rest"
	"DoctorWho/internal/pkg/Bot"
	"DoctorWho/internal/pkg/middleware"
	"DoctorWho/internal/repository"
	"DoctorWho/internal/usecase"

	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	//r := gin.Default()
	//r.Use(gin.Recovery())
	//r.Use
	//apiGroup := r.Group("/api")
	//{
	//
	//	apiGroup.POST("/signup", controller.SignUp)
	//	apiGroup.POST("/login", controller.Login)
	//	apiGroup.POST("/verification", controller.Verification)
	//	apiGroup.POST("/logout", controller.Logout)
	//}

	instance := configs.Configuration()
	postgres, err := configs.NewPostgresConfig(*instance)
	bot, err := configs.NewBotConfig(*instance)
	NewBot := Bot.NewBot(bot)
	if err != nil {
		panic("no connection")
	}
	repo := repository.NewRepo(postgres, NewBot)
	service := usecase.NewUserUsecase(repo, NewBot)
	controller := rest.NewController(service)
	r := mux.NewRouter()

	r.HandleFunc("/signup", controller.SignUp)
	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/verification", controller.Verification)
	r.Use(middleware.AuthMiddleWare)
	r.HandleFunc("/logout", controller.Logout)
	NewBot.SendNotification("listening on :5005")
	server := &http.Server{Addr: ":5005", Handler: r}
	server.ListenAndServe()
}
