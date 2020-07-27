package storage

import (
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
	logger *zap.Logger
}

func NewPsqlStorage(cfg *Config, logger *zap.Logger) (Storage, error) {
	opts, err := configToPgOptions(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "can not create database")
	}

	d := &dataBase{
		db:     pg.Connect(opts),
		logger: logger.Named("psqlstorage"),
	}

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
