package storage

import (
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/model"
)

func (d *dataBase) AddUser(user *model.User) error {
	dtoModel, err := modelUserToDtoUser(user)
	if err != nil {
		return err
	}

	if _, err = d.db.Model(dtoModel).Insert(); err != nil {
		return sqlErrorToError(errors.Wrap(err, "query processing error"))
	}

	return nil
}
