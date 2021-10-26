package main

import (
	"jmind-test/src/config"
	"jmind-test/src/routes"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting the jmind-test-api service...")
	config.ServerCtx = config.InitServerContext()
	router := routes.Router()
	log.Fatal(http.ListenAndServe(":5000", router))
}
