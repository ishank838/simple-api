package ping

import (
	"encoding/json"
	"net/http"

	"github.com/ishank838/simple-api/logger"
	"github.com/jmoiron/sqlx"
)

type service struct {
	svc pingService
}

type pingResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func PingHandler(svc pingService) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		message, err := svc.PingDb()
		if err != nil {
			writeError(rw, http.StatusInternalServerError, "ping", err)
		}
		if err := json.NewEncoder(rw).Encode(message); err != nil {
			logger.ErrorOf("[Ping] error writing response: %v", err)
		}
	}
}

func writeError(w http.ResponseWriter, status int, message string, err error) {
	pingResponse := struct{ Error string }{Error: message}
	w.WriteHeader(status)
	logger.ErrorOf("ping failiure %s with error %v", message, err)
	if err = json.NewEncoder(w).Encode(pingResponse); err != nil {
		logger.ErrorOf("ping error while encoding response: %v", err)
	}
}

func NewService(db *sqlx.DB) pingService {
	return pingService{
		store: pingStore{db},
	}
}
