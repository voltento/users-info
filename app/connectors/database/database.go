package database

import (
	"context"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/modules/consumer"
	"github.com/voltento/users-info/app/modules/storage"
	"go.uber.org/zap"
)

const (
	tableNameUsersInfo    = "users"
	tableColumnNameUserId = "user_id"
)

// Implements modules.Storage
type dataBase struct {
	consumer consumer.Consumer
	db       *pg.DB
	logger   *zap.SugaredLogger
	cfg      *Config
}

func NewDataBase(cfg *Config, consumer consumer.Consumer, logger *zap.SugaredLogger) (storage.Storage, error) {
	opts, err := configToPgOptions(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "can not create database")
	}

	d := &dataBase{
		db:       pg.Connect(opts),
		logger:   logger.Named("psqlstorage"),
		cfg:      cfg,
		consumer: consumer,
	}

	d.logDbConnection()

	err = d.Check(context.TODO())
	if err != nil {
		return nil, errors.Wrap(err, "can not create database")
	}

	return d, nil
}

func (d *dataBase) Stop() error {
	return d.db.Close()
}

func (d *dataBase) logDbConnection() {
	opts := d.db.Options()
	d.logger.Infof("db configuration: address '%v' user '%v'", opts.Addr, opts.User)
}
