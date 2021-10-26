package config

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log"
)

var ServerCtx *ServerContext

type ServerContext struct {
	Db *pg.DB
}

func InitServerContext() *ServerContext {
	Db := ConnectToDb()
	if err := Db.Ping(context.Background()); err != nil {
		log.Fatalf("Error connecting to DB: %s.", err.Error())
	} else {
		log.Print("DB connected successfully.")

	}

	return &ServerContext{
		Db: Db,
	}
}
