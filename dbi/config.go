package dbi

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Driver   string `default:"postgres"`
	Host     string `required:"true"`
	Port     int    `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
}

func (conf Config) dbUrl() string {
	return fmt.Sprintf("host=%s port=%v user=%s password=%s", conf.Host, conf.Port, conf.User, conf.Password)
}

func NewDb(config Config) (*sqlx.DB, error) {

	db, err := sqlx.Open(config.Driver, config.dbUrl())
	if err != nil {
		return nil, fmt.Errorf("failed to load databse config %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping server %v", err)
	}

	return db, nil
}
