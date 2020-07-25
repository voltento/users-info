package main

import (
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/config"
	db "github.com/voltento/users-info/app/connectors/psql_connector"
	"github.com/voltento/users-info/app/logger"
	"github.com/voltento/users-info/app/service"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"log"
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
	serv.Stop()
	err = errors.Wrap(err, "server finished")
}

func di() (error, *dig.Container) {
	di := dig.New()
	err := di.Provide(db.NewPsqlStorage)
	if err != nil {
		return err, nil
	}

	err = di.Provide(config.NewDefaultConfigs)
	if err != nil {
		return err, nil
	}

	err = di.Provide(func(cfg *logger.Config) (*zap.Logger, error) {
		return logger.NewLogger(cfg)
	})
	if err != nil {
		return err, nil
	}

	err = di.Provide(func(logger *zap.Logger) *zap.SugaredLogger {
		return logger.Sugar()
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
