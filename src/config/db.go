package config

import (
	"github.com/go-pg/pg/v10"
	"os"
)

var Addr = os.Getenv("DB_ADDR")
var User = os.Getenv("DB_USER")
var Password = os.Getenv("DB_PASSWORD")
var DBname = os.Getenv("DB_NAME")

var (
	Db *pg.DB
)

func ConnectToDb() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     Addr,
		User:     User,
		Password: Password,
		Database: DBname,
	})
}
