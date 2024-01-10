package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"testDeployment/internal/pkg/Bot"
	"testDeployment/internal/usecase"
)

type controller struct {
	usecase usecase.Usecase
	bot     Bot.Bot
}

func NewController(usecase usecase.Usecase, bot Bot.Bot) *controller {
	return &controller{usecase: usecase,
		bot: bot,
	}
}

func (c controller) GetAll(w http.ResponseWriter, r *http.Request) {
	user := c.usecase.GetAll()
	log.Println(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}
