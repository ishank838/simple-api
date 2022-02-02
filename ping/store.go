package ping

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type pingStore struct {
	db *sqlx.DB
}

func (store *pingStore) pindDB() (string, error) {
	if err := store.db.Ping(); err != nil {
		return "", fmt.Errorf("failed to ping db %v", err)
	}
	return "pong", nil
}
