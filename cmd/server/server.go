package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ishank838/simple-api/config"
	"github.com/ishank838/simple-api/dbi"
	"github.com/ishank838/simple-api/ping"
)

func Server(app config.Application) (*mux.Router, error) {
	m := mux.NewRouter()

	db, err := dbi.NewDb(app.Db)
	if err != nil {
		return nil, err
	}

	pingService := ping.NewService(db)
	m.HandleFunc("/ping", ping.PingHandler(pingService)).Methods(http.MethodGet)
	return m, nil
}
