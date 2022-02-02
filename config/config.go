package config

import (
	"fmt"

	"github.com/ishank838/simple-api/dbi"
	"github.com/ishank838/simple-api/logger"
	"github.com/kelseyhightower/envconfig"
)

var app Application

type serverConfig struct {
	Port     string `required:"true" default:"8080"`
	Host     string `required:"true"`
	Protocol string `default:"http"`
}

type Application struct {
	server serverConfig
	Db     dbi.Config
}

func Address() string {
	return fmt.Sprintf("%s:%v", app.server.Host, app.server.Port)
}

func MustLoad() Application {
	err := envconfig.Process("S", &app.server)
	if err != nil {
		logger.Fatal("failed to load server config", err)
	}

	err = envconfig.Process("DB", &app.Db)
	if err != nil {
		logger.Fatal("failed to load db config", err)
	}
	return app
}
