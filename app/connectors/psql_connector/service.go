package psql_connector

import (
	"context"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/connectors"
	"github.com/voltento/users-info/app/model"
)

const (
	tableNameUsersInfo = "users"
)

type dataBase struct {
	db *pg.DB
}

func NewDatabase(cfg *Config) (error, connectors.Storage) {
	err, opts := configToPgOptions(cfg)
	if err != nil {
		return errors.Wrap(err, "can not create database"), nil
	}

	d := &dataBase{
		db: pg.Connect(opts),
	}

	err = d.HealthCheck()
	if err != nil {
		return errors.Wrap(err, "can not create database"), nil
	}

	return nil, d
}

func (d *dataBase) Users() (error, []model.User) {
	// todo: add log
	err, dtoUsers := d.GetUsers()
	if err != nil {
		return errors.Wrap(err, "cant get users"), nil
	}
	users := make([]model.User, 0, len(dtoUsers))
	for _, u := range dtoUsers {
		users = append(users, dtoUserToModelUser(u))
	}
	return nil, users
}

func (d *dataBase) HealthCheck() error {
	var n int
	_, err := d.db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func (d *dataBase) GetUsers() (error, []*User) {
	return d.GetUsersWithCtx(nil)
}

func (d *dataBase) GetUsersWithCtx(ctx context.Context) (error, []*User) {
	var users []*User
	pg.Scan("")
	err := d.db.WithContext(ctx).Model().Table(tableNameUsersInfo).Select(&users)

	return err, users
}
