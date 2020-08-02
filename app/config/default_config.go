package config

import (
	"github.com/voltento/users-info/app/connectors/database"
	"github.com/voltento/users-info/app/logger"
	"github.com/voltento/users-info/app/service"
	"net"
)

func NewDefaultConfig() *Config {
	dbConfig := &database.Config{
		User:             "users-info",
		Password:         "users-info",
		Database:         "users-info",
		LimitGetEntities: 500,
	}
	loggerConfig := &logger.Config{
		Level: "debug",
	}

	serviceConfig := &service.Config{
		Address:     net.JoinHostPort("localhost", "8181"),
		LogGinGonic: false,
	}

	return &Config{dbConfig, loggerConfig, serviceConfig}
}
