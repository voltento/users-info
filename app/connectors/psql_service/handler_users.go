package psql_service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/model"
)

func (d *dataBase) Users() ([]model.User, error) {
	dtoUsers, err := d.GetUsers()
	if err != nil {
		d.logger.Error("failed")
		return nil, errors.Wrap(err, "cant get users")
	}
	d.logger.Sugar().Debugf("get %v users from database", len(dtoUsers))
	users := make([]model.User, 0, len(dtoUsers))
	for _, u := range dtoUsers {
		users = append(users, dtoUserToModelUser(u))
	}
	return users, nil
}

func (d *dataBase) GetUsers() ([]*User, error) {
	return d.GetUsersWithCtx(context.TODO())
}

func (d *dataBase) GetUsersWithCtx(ctx context.Context) ([]*User, error) {
	var users []*User
	err := d.db.WithContext(ctx).Model().Table(tableNameUsersInfo).Select(&users)

	return users, err
}
