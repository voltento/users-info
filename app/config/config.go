package config

import (
	"flag"
	db "github.com/voltento/users-info/app/connectors/psql_service"
	"github.com/voltento/users-info/app/logger"
	"github.com/voltento/users-info/app/service"
	"log"
)

type Config struct {
	Database *db.Config
	Logger   *logger.Config
	Service  *service.Config
}

func NewConfig() (*Config, error) {
	path := flag.String("config", "", "help message for flagname")
	flag.Parse()
	if path == nil || len(*path) == 0 {
		log.Printf("no config path. use default config")
		return NewDefaultConfig(), nil
	}

	log.Printf("load config from file '%v'", *path)
	return NewConfigFromFile(*path)
}

func GetConfigs(cfg *Config) (*db.Config, *logger.Config, *service.Config) {
	return cfg.Database, cfg.Logger, cfg.Service
}
