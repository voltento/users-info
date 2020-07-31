package storage

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/fault"
	"github.com/voltento/users-info/app/model"
	"strconv"
	"strings"
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

	if len(cfg.Addr) != 0 {
		opts.Addr = cfg.Addr
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
		LastName:    dtoUser.LastName,
		Email:       dtoUser.Email,
		CountryCode: dtoUser.CountryCode,
	}
}

func modelUserToDtoUser(user *model.User) (*User, error) {
	u := &User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		CountryCode: user.CountryCode,
	}

	if len(user.UserId) > 0 {
		id, err := strconv.Atoi(user.UserId)
		if err != nil {
			err = errors.Wrap(err, fmt.Sprintf("can not convert userId '%v' to the dto user id", user.UserId))
			return nil, fault.NewBadRequest(err.Error())
		}

		u.UserId = id
	}

	return u, nil
}

func sqlErrorToError(err error) error {
	if strings.Contains(err.Error(), "violates not-null constraint") {
		return fault.NewBadRequest(err.Error())
	}

	if strings.Contains(err.Error(), "duplicate key value violates") {
		return fault.NewBadRequest(err.Error())
	}
	return err
}
