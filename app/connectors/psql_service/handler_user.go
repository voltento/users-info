package psql_service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/model"
)

func (d *dataBase) User(userId string) (*model.User, error) {
	dtoUser, err := d.GetUser(userId)
	if err != nil {
		d.logger.Error("failed")
		return nil, errors.Wrap(err, "cant get users")
	}
	d.logger.Sugar().Debugf("get user from database")

	user := dtoUserToModelUser(dtoUser)
	return &user, nil
}

func (d *dataBase) GetUser(userId string) (*User, error) {
	return d.GetUserWithCtx(userId, context.TODO())
}

func (d *dataBase) GetUserWithCtx(userId string, ctx context.Context) (*User, error) {
	var users []*User
	err := d.db.WithContext(ctx).Model().Table(tableNameUsersInfo).Where(tableColumnNameUserId+" = ?", userId).Select(&users)

	if len(users) == 0 {
		return nil, errors.New(fmt.Sprintf("can not find any users by id = %v", userId))
	}

	if len(users) > 0 {
		d.logger.Named("GetUserWithCtx").Error(fmt.Sprintf("got %v records, expected 1", len(users)))
	}

	return users[0], err
}
