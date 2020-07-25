package main

import (
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/connectors"
	db "github.com/voltento/users-info/app/connectors/psql_connector"
	"github.com/voltento/users-info/app/service"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"log"
	"net"
)

func main() {
	err, di := di()
	if err != nil {
		log.Fatal(err.Error())
	}

	var serv *service.Service
	var logger *zap.Logger
	defer logger.Sync()
	err = di.Invoke(func(s *service.Service, l *zap.Logger) {
		serv = s
		logger = l
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = serv.Run()
	err = errors.Wrap(err, "server finished")
}

func di() (error, *dig.Container) {
	di := dig.New()
	err := di.Provide(db.NewPsqlStorage)
	if err != nil {
		return err, nil
	}

	err = di.Provide(func() (*zap.Logger, error) {
		return zap.NewProduction()
	})
	if err != nil {
		return err, nil
	}

	err = di.Provide(func() *db.Config {
		return &db.Config{
			User:     "users-info",
			Password: "users-info",
			Database: "users-info",
		}
	})
	if err != nil {
		return err, nil
	}

	err = di.Provide(func(s connectors.Storage) service.Config {
		return service.Config{
			Addr:        net.JoinHostPort("localhost", "8181"),
			Storage:     s,
			LogGinGonic: false,
		}
	})
	if err != nil {
		return err, nil
	}

	err = di.Provide(service.NewService)
	if err != nil {
		return err, nil
	}

	return err, di
}
