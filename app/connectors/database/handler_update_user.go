package database

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/fault"
	"github.com/voltento/users-info/app/model"
)

func (d *dataBase) UpdateUser(user *model.User) error {
	dtoModel, err := modelUserToDtoUser(user)
	if err != nil {
		return err
	}

	r, err := d.db.Model(dtoModel).WherePK().Update()

	if err != nil {
		return sqlErrorToError(errors.Wrap(err, "query processing error"))
	}

	if r.RowsAffected() == 0 {
		return fault.NewNotFound(fmt.Sprintf("can not find any user by id %v", user.UserId))
	}

	return nil
}
