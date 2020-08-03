package database

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/fault"
	"github.com/voltento/users-info/app/model"
)

func (d *dataBase) DropUser(userId string) error {
	ormUser, err := modelUserToDtoUser(&model.User{
		UserId: userId,
	})

	if err != nil {
		return fault.NewBadRequest(fmt.Sprintf("can not convert %v to user id", userId))
	}

	r, err := d.db.Model(ormUser).Table(tableNameUsersInfo).WherePK().Delete()
	if err != nil {
		return sqlErrorToError(errors.Wrap(err, "query processing error"))
	}

	if r.RowsAffected() == 0 {
		return fault.NewNotFound(fmt.Sprintf("can not find any user by id %v", userId))
	}

	d.logger.Infof("dropped user. user_id '%v'", userId)

	return nil
}
