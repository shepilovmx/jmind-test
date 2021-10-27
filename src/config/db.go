package config

import (
	"github.com/go-pg/pg/v10"
	"os"
)

func ConnectToDb() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_ADDR"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})
}
