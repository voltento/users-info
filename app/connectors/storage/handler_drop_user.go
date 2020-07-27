package storage

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/cerrors"
	"github.com/voltento/users-info/app/model"
)

func (d *dataBase) DropUser(userId string) error {
	ormUser, err := modelUserToDtoUser(&model.User{
		UserId: userId,
	})

	if err != nil {
		return errors.New(fmt.Sprintf("can not convert %v to user id", userId))
	}

	r, err := d.db.Model(ormUser).Table(tableNameUsersInfo).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, "query processing error")
	}

	if r.RowsAffected() == 0 {
		return cerrors.NewErrorNotFound(fmt.Sprintf("can not find any user by id %v", userId))
	}

	return nil
}
