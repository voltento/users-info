package storage

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
}

func (d *dataBase) Check(_ context.Context) error {
	return d.HealthCheck()
}

func NewPsqlStorage(cfg *Config, logger *zap.SugaredLogger) (Storage, error) {
	opts, err := configToPgOptions(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "can not create database")
	}

	d := &dataBase{
		db:     pg.Connect(opts),
		logger: logger.Named("psqlstorage"),
	}

	d.logDbConnection()

	err = d.HealthCheck()
	if err != nil {
		return nil, errors.Wrap(err, "can not create database")
	}

	return d, nil
}

func (d *dataBase) HealthCheck() error {
	var n int
	_, err := d.db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func (d *dataBase) Stop() error {
	return d.db.Close()
}

func (d *dataBase) logDbConnection() {
	opts := d.db.Options()
	d.logger.Infof("db configuration: address '%v' user '%v'", opts.Addr, opts.User)
}
