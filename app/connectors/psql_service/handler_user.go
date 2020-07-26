package psql_service

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/model"
)

func (d *dataBase) User(modelUser *model.User) (*model.User, error) {
	dtoModel, err := modelUserToDtoUser(modelUser)
	if err != nil {
		return nil, err
	}

	dtoUser, err := d.GetUser(dtoModel)
	if err != nil {
		d.logger.Error("failed")
		return nil, errors.Wrap(err, "cant get users")
	}
	d.logger.Sugar().Debugf("get user from database")

	user := dtoUserToModelUser(dtoUser)
	return &user, nil
}

func (d *dataBase) GetUser(ormUser *User) (*User, error) {
	return d.GetUserWithCtx(ormUser, context.TODO())
}

func (d *dataBase) GetUserWithCtx(ormUser *User, ctx context.Context) (*User, error) {
	var users []*User
	q := d.db.WithContext(ctx).Model().Table(tableNameUsersInfo)
	q = buildUserWhereEqual(ormUser, q)
	err := q.Select(&users)

	if len(users) == 0 {
		return nil, errors.New(fmt.Sprintf("can not find any users by model = %v", ormUser))
	}

	if len(users) > 1 {
		d.logger.Named("GetUserWithCtx").Error(fmt.Sprintf("got %v records, expected 1", len(users)))
	}

	return users[0], err
}
