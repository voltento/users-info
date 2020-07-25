package config

import (
	db "github.com/voltento/users-info/app/connectors/psql_connector"
	"github.com/voltento/users-info/app/logger"
	"github.com/voltento/users-info/app/service"
	"net"
)

func NewDefaultConfigs() (*logger.Config, *db.Config, *service.Config) {
	dbConfig := &db.Config{
		User:     "users-info",
		Password: "users-info",
		Database: "users-info",
	}
	loggerConfig := &logger.Config{
		Level: "debug",
	}

	serviceConfig := &service.Config{
		Addr:        net.JoinHostPort("localhost", "8181"),
		LogGinGonic: false,
	}

	return loggerConfig, dbConfig, serviceConfig
}
