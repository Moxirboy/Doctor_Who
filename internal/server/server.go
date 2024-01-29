package server

import (
	"fmt"
	"os"
	configs "testDeployment/internal/common/config"
	"testDeployment/internal/delivery"
	request "testDeployment/internal/delivery/http"
		"testDeployment/pkg/Bot"
		"testDeployment/internal/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)
type Server struct{
	cfg *configs.Config
}
func NewServer(
	cfg *configs.Config,
	) *Server{
	return &Server{
		cfg: cfg,
	}
}
func (s Server) Run() error{

	r:=gin.New()
	store := cookie.NewStore([]byte("curifyDoctorWho"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   30 * 24 * 60 * 60, // Session expires in 30 days (in seconds)
		HttpOnly: true,
		Secure:   true, // Set Secure to true for HTTPS-only
	})
	r.Use(sessions.Sessions("mysession", store))
	r.Use(gin.Recovery())
	token := "6509175022:AAFUZtIJBUJuzW_WuRfbz08AQcYWbL6YYEI"
	bot, err := configs.BotConfi(token)
	NewBot := Bot.NewBot(bot)
	if err != nil {
		NewBot.SendErrorNotification(err)
		return err
		
	}

	httpClient := request.NewCustomHTTPClient()
	jsonRequester := request.NewCustomJSONRequester(httpClient)
	pg, err := configs.NewPostgresConfig()
	if err != nil {
		NewBot.SendErrorNotification(err)
		return err
	}
	uc:=usecase.New(pg,NewBot)
	delivery.SetUp(r,uc,NewBot,*jsonRequester)
	Port:= os.Getenv("PORT")
	if Port == "" {
		Port = "3000"
	}
	NewBot.SendNotification("Runnung on : "+Port)
	return r.Run(fmt.Sprintf(":%s",Port ))
}