package main

import (
	"jmind-test/src/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("ETHERSCAN_API_URL", "https://api.etherscan.io/api")
	os.Setenv("ETHERSCAN_API_KEY", "YourApiKeyToken")
	log.Print("Starting the jmind-test-api service...")

	router := routes.Router()
	log.Fatal(http.ListenAndServe(":5000", router))
}
