package psql_connector

import (
	"errors"
	"github.com/go-pg/pg"
	"github.com/voltento/users-info/app/model"
)

// TODO: add tests
func configToPgOptions(cfg *Config) (error, *pg.Options) {
	if cfg == nil {
		return errors.New("empty value provided"), nil
	}

	opts := &pg.Options{
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
	}

	if len(opts.User) == 0 {
		return errors.New("empty user provided"), nil
	}

	if len(opts.Password) == 0 {
		return errors.New("empty password provided"), nil
	}

	if len(opts.Database) == 0 {
		return errors.New("empty database provided"), nil
	}

	return nil, opts
}

// TODO: add tests
func dtoUserToModelUser(dtoUser *User) model.User {
	return model.User{
		UserId:      dtoUser.UserId,
		FirstName:   dtoUser.FirstName,
		SecondName:  dtoUser.LastName,
		Email:       dtoUser.Email,
		CountryCode: dtoUser.CountryCode,
	}
}
