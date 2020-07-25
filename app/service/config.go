package service

import "github.com/voltento/users-info/app/connectors"

type Config struct {
	Storage     connectors.Storage
	Addr        string
	LogGinGonic bool
}
