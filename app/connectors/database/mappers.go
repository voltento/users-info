package database

import (
	"fmt"
	"github.com/badoux/checkmail"
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

	if len(user.Email) > 0 {
		if err := checkmail.ValidateFormat(user.Email); err != nil {
			return nil, fault.NewBadRequest(errors.Wrap(err, "invalid email").Error())
		}
	}

	return u, nil
}

func sqlErrorToError(err error) error {
	badRequestStrings := []string{"violates not-null constraint", "duplicate key value violates"}

	for _, s := range badRequestStrings {
		if strings.Contains(err.Error(), s) {
			return fault.NewBadRequest(err.Error())
		}
	}
	return err
}
