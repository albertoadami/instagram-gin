package database

import (
	"fmt"

	"github.com/albertoadami/instagram-gin/internal/configuration"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(config *configuration.DatabaseConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name,
	)
	return sqlx.Connect("postgres", dsn)
}
