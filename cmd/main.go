package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	configs "testDeployment/internal/common/config"
	"testDeployment/internal/delivery/rest"
	"testDeployment/internal/pkg/Bot"
	"testDeployment/internal/pkg/middleware"
	"testDeployment/internal/repository"
	"testDeployment/internal/usecase"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	store := cookie.NewStore([]byte("curifyDoctorWho"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   30 * 24 * 60 * 60, // Session expires in 30 days (in seconds)
		HttpOnly: true,
		Secure:   true, // Set Secure to true for HTTPS-only
	})
	r.Use(sessions.Sessions("mysession", store))
	r.Use(gin.Recovery())
	token := "6497806085:AAE7RRULTJARZ5YTP163We6ri4mnSK1OHBk"
	bot, err := configs.BotConfi(token)
	NewBot := Bot.NewBot(bot)
	if err != nil {
		NewBot.SendErrorNotification(err)
		return
	}
	config, err := configs.NewPostgresConfig()
	if err != nil {
		NewBot.SendErrorNotification(err)
		return
	}
	repo := repository.NewRepo(config, NewBot)
	service := usecase.NewUserUsecase(repo, NewBot)
	controller := rest.NewController(service, NewBot)
	save := r.Group("/save")
	{
		save.GET("/", func(c *gin.Context) {
			c.String(200, "Hello from save")
		})
		program := save.Group("/program")
		{

			program.GET("/", controller.ProgramHandler)
			program.POST("/upload", controller.NewProgram)
		}
		exercise := save.Group("/exercise")
		{
			exercise.GET("/", controller.ExerciseHandler)
			exercise.POST("/upload", controller.Newxercise)
		}
		drugs := save.Group("/drugs")
		{
			drugs.GET("/", controller.DrugIndexHandler)
			drugs.POST("/upload", controller.DrugUploadHandler)
		}
	}
	v1Group := r.Group("/v1")
	{
		v1Group.GET("/hello", func(c *gin.Context) {
			c.String(200, "Hello, World!")
		})
		v1Group.POST("/signup", controller.SignUp)
		v1Group.POST("/login", controller.Login)

		v1Group.POST("/verification", controller.Verification)
		dash := v1Group.Group("/dashboard")
		{
			dash.GET("/", func(c *gin.Context) {
				c.String(200, "Hello from dashboard")
			})
			dash.GET("/searchDrug", controller.SearchDrug)
			dash.GET("/getOneDrugById", controller.GetDrug)
			dash.GET("/getdrug",controller.GetAllDrug)
			middle := dash.Group("/middle")
			middle.Use(middleware.AuthMiddleware())
			{
				middle.GET("/getexerciseforweightloss", controller.GetProgramForWeightLoss)
				middle.GET("/getexerciseforstress", controller.GetProgramForStress)
				middle.POST("/markasdone", controller.MarkAsDone)
				middle.POST("/updateuserinfo", controller.UpdateUserInfo)
				middle.GET("/showUserInfo", controller.ShowUserInfo)
				middle.GET("/logout", controller.Logout)
				middle.GET("/deleteAccount", controller.DeleteAccount)
			}
			{
			}
			dash.POST("/fillUserInfo", controller.FillUserInfo)
			dash.GET("/getstressWorkprogress", controller.GetProgressStreesWork)
			dash.GET("/getweightlossprogress", controller.GetProgramForWeightLoss)
		}

	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	NewBot.SendNotification("listening on :" + port)
	log.Println("listening on :" + port)
	server := &http.Server{
		Addr:         "0.0.0.0:" + port,
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
		NewBot.SendErrorNotification(err)

	} else {
		NewBot.SendNotification("Server has been gracefully stopped.")

	}

}
