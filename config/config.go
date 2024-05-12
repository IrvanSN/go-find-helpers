package config

import (
	"github.com/irvansn/go-find-helpers/drivers/postgresql"
	"os"
)

func InitConfigPostgresql() postgresql.Config {
	return postgresql.Config{
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
	}
}
