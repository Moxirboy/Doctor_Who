package Bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
)

type bot struct {
	*tgbotapi.BotAPI
}
type Bot interface {
	SendErrorNotification(err error)
	SendNotification(mess string)
}

func NewBot(botAPI *tgbotapi.BotAPI) Bot {
	return &bot{BotAPI: botAPI}
}

const chatID = int64(-1002129341182)

func (b bot) SendErrorNotification(err error) {
	// Replace "USER_CHAT_ID" with the actual user's chat ID to send the notification

	message := time.Now().Format("2006/02/31  15:04:05\n") + "Error occurred: "
	msg := tgbotapi.NewMessage(chatID, message+err.Error())
	_, err = b.Send(msg)
	if err != nil {
		log.Printf("Error sending notification to user: %v", err)
	}

}
func (b bot) SendNotification(mess string) {
	// Replace "USER_CHAT_ID" with the actual user's chat ID to send the notification

	message := time.Now().Format("2006/02/31  15:04:05 \nmessage: ") + mess
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := b.Send(msg)
	if err != nil {
		log.Printf("Error sending notification to user: %v", err)
	}
}
