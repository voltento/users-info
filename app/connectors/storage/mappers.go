package storage

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/model"
	"strconv"
)

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

func dtoUserToModelUser(dtoUser *User) model.User {
	return model.User{
		UserId:      strconv.Itoa(dtoUser.UserId),
		FirstName:   dtoUser.FirstName,
		SecondName:  dtoUser.LastName,
		Email:       dtoUser.Email,
		CountryCode: dtoUser.CountryCode,
	}
}

func modelUserToDtoUser(user *model.User) (*User, error) {
	u := &User{
		FirstName:   user.FirstName,
		LastName:    user.SecondName,
		Email:       user.Email,
		CountryCode: user.CountryCode,
	}

	if len(user.UserId) > 0 {
		id, err := strconv.Atoi(user.UserId)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("can not convert userId '%v' to the dto user id", user.UserId))
		}

		u.UserId = id
	}

	return u, nil
}
