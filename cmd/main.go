package main

import (
	"log"
	"net/http"

	"github.com/ishank838/simple-api/cmd/server"
	"github.com/ishank838/simple-api/config"
	"github.com/ishank838/simple-api/logger"
)

func main() {
	app := config.MustLoad()

	r, err := server.Server(app)

	if err != nil {
		logger.Fatal("[main]failed to initialise server", err)
	}

	address := config.Address()
	logger.InfoOf("[main]Serving at address: %s", address)
	if err = http.ListenAndServe(address, r); err != nil {
		log.Fatal("[main]failed to start server", err)
	}
}
