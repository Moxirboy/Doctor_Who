package configs

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func NewPostgresConfig() (*sql.DB, error) {
	instance, err := sql.Open("pgx", "postgresql://postgres:6Ea2Df4*fg15eF1cFd33bA555-dF*3DA@viaduct.proxy.rlwy.net:25923/railway")
	if err != nil {
		panic(err)
	}
	err = instance.Ping()
	if err != nil {
		panic(err)
	}

	return instance, nil
}



func NewSQLiteConfig(sqliteFilePath string) (*sql.DB, error) {
	instance, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&mode=rwc", sqliteFilePath))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = instance.Ping()
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return instance, nil
}
func BotConfi(as string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(as)
	if err != nil {
		log.Println(err)
	}
	return bot, nil
}


func NewBotConfig(cfg Config) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Println(err)
	}
	return bot, nil
}
