package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewPostgresConfig(cfg Config) (*sql.DB, error) {
	psqlString := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SslMode)

	instance, err := sql.Open("pgx", psqlString)
	if err != nil {
		panic(err)
	}
	err = instance.Ping()
	if err != nil {
		panic(err)
	}

	return instance, nil
}
