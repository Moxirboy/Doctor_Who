package main

import (
	configs "DoctorWho/internal/common/config"
	"DoctorWho/internal/delivery/rest"
	"DoctorWho/internal/pkg/Bot"
	"DoctorWho/internal/pkg/middleware"
	"context"
	"log"
	"time"

	"DoctorWho/internal/repository"
	"DoctorWho/internal/usecase"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(gin.Recovery())
	instance := configs.Configuration()
	postgres, err := configs.NewPostgresConfig(*instance)
	bot, err := configs.NewBotConfig(*instance)
	NewBot := Bot.NewBot(bot)
	if err != nil {
		NewBot.SendErrorNotification(err)
		return
	}
	repo := repository.NewRepo(postgres, NewBot)
	service := usecase.NewUserUsecase(repo, NewBot)
	controller := rest.NewController(service)

	v1Group := r.Group("/v1")
	{

		v1Group.GET("/signup", controller.SignUp)
		v1Group.GET("/login", controller.Login)
		v1Group.GET("/verification", controller.Verification)
		dash:=v1Group.Group("dashboard")
		dash.Use(middleware.AuthMiddleware())
		{
			dash.GET("/logout", controller.Logout)
		}
		
	}
	NewBot.SendNotification("listening on :5005")
	log.Println("listening on :5005")
	server := &http.Server{
		Addr:         "0.0.0.0:5005",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			NewBot.SendErrorNotification(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	NewBot.SendNotification("Server is shutting down...")
	

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		NewBot.SendErrorNotification( err)
	
	} else {
		NewBot.SendNotification("Server has been gracefully stopped.")
		
	}

}
