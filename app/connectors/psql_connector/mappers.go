package psql_connector

import (
	"errors"
	"github.com/go-pg/pg"
	"github.com/voltento/users-info/app/model"
)

// TODO: add tests
func configToPgOptions(cfg *Config) (*pg.Options, error) {
	if cfg == nil {
		return nil, errors.New("empty value provided")
	}

	opts := &pg.Options{
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
	}

	if len(opts.User) == 0 {
		return nil, errors.New("empty user provided")
	}

	if len(opts.Password) == 0 {
		return nil, errors.New("empty password provided")
	}

	if len(opts.Database) == 0 {
		return nil, errors.New("empty database provided")
	}

	return opts, nil
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
