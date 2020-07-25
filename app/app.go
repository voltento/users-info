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
	di, err := di()
	if err != nil {
		log.Fatal(err.Error())
	}

	var serv *service.Service
	var zapLogger *zap.Logger

	err = di.Invoke(func(s *service.Service, l *zap.Logger) {
		serv = s
		zapLogger = l
	})

	defer func() {
		err := zapLogger.Sync()
		if err != nil {
			log.Printf("error during sync logger: %v", err.Error())
		}
	}()

	if err != nil {
		log.Fatal(err.Error())
	}

	err = serv.Run()
	serv.Stop()
	err = errors.Wrap(err, "server finished")
}

func di() (*dig.Container, error) {
	di := dig.New()
	err := di.Provide(db.NewPsqlStorage)
	if err != nil {
		return nil, err
	}

	err = di.Provide(config.NewDefaultConfigs)
	if err != nil {
		return nil, err
	}

	err = di.Provide(func(cfg *logger.Config) (*zap.Logger, error) {
		return logger.NewLogger(cfg)
	})
	if err != nil {
		return nil, err
	}

	err = di.Provide(func(logger *zap.Logger) *zap.SugaredLogger {
		return logger.Sugar()
	})
	if err != nil {
		return nil, err
	}

	err = di.Provide(service.NewService)
	if err != nil {
		return nil, err
	}

	return di, err
}
