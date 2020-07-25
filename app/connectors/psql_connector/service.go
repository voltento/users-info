package psql_connector

import (
	"context"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/connectors"
	"github.com/voltento/users-info/app/model"
	"go.uber.org/zap"
)

const (
	tableNameUsersInfo = "users"
)

type dataBase struct {
	db     *pg.DB
	logger *zap.Logger
}

func NewPsqlStorage(cfg *Config, logger *zap.Logger) (connectors.Storage, error) {
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

func (d *dataBase) Users() ([]model.User, error) {
	dtoUsers, err := d.GetUsers()
	if err != nil {
		d.logger.Error("failed")
		return nil, errors.Wrap(err, "cant get users")
	}
	d.logger.Sugar().Debugf("get %v users from database", len(dtoUsers))
	users := make([]model.User, 0, len(dtoUsers))
	for _, u := range dtoUsers {
		users = append(users, dtoUserToModelUser(u))
	}
	return users, nil
}

func (d *dataBase) HealthCheck() error {
	var n int
	_, err := d.db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func (d *dataBase) GetUsers() ([]*User, error) {
	return d.GetUsersWithCtx(nil)
}

func (d *dataBase) GetUsersWithCtx(ctx context.Context) ([]*User, error) {
	var users []*User
	pg.Scan("")
	err := d.db.WithContext(ctx).Model().Table(tableNameUsersInfo).Select(&users)

	return users, err
}

func (d *dataBase) Stop() error {
	return d.db.Close()
}
