package database

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/model"
)

func (d *dataBase) User(userId string) (*model.User, error) {
	dtoModel, err := modelUserToDtoUser(&model.User{
		UserId: userId,
	})

	if err != nil {
		return nil, err
	}

	var users []*User
	users, err = d.GetUsers(dtoModel)
	if err != nil {
		return nil, sqlErrorToError(err)
	}

	if len(users) == 0 {
		return nil, errors.New(fmt.Sprintf("can not find any users by id = %v", userId))
	}

	if len(users) > 1 {
		d.logger.Named("GetUserWithCtx").Error(fmt.Sprintf("got %v records, expected 1", len(users)))
	}

	user := dtoUserToModelUser(users[0])
	return &user, nil
}
