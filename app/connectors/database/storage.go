package database

import (
	"context"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	tableNameUsersInfo    = "users"
	tableColumnNameUserId = "user_id"
)

type dataBase struct {
	db     *pg.DB
	logger *zap.SugaredLogger
	cfg    *Config
}

func NewDataBase(cfg *Config, logger *zap.SugaredLogger) (Storage, error) {
	opts, err := configToPgOptions(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "can not create database")
	}

	d := &dataBase{
		db:     pg.Connect(opts),
		logger: logger.Named("psqlstorage"),
		cfg:    cfg,
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
