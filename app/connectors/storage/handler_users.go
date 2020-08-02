package storage

import (
	"context"
	"github.com/pkg/errors"
	"github.com/voltento/users-info/app/model"
)

func (d *dataBase) Users(modelUser *model.User) ([]model.User, error) {
	dtoModel, err := modelUserToDtoUser(modelUser)
	if err != nil {
		return nil, err
	}

	dtoUsers, err := d.GetUsers(dtoModel)
	if err != nil {
		d.logger.Error("failed")
		return nil, sqlErrorToError(errors.Wrap(err, "cant get users"))
	}
	d.logger.Debugf("get %v users from database", len(dtoUsers))
	users := make([]model.User, 0, len(dtoUsers))
	for _, u := range dtoUsers {
		users = append(users, dtoUserToModelUser(u))
	}
	return users, nil
}

func (d *dataBase) GetUsers(ormUser *User) ([]*User, error) {
	return d.GetUsersWithCtx(ormUser, context.TODO())
}

func (d *dataBase) GetUsersWithCtx(ormUser *User, ctx context.Context) ([]*User, error) {
	var users []*User
	q := d.db.WithContext(ctx).Model().Table(tableNameUsersInfo).Limit(d.cfg.LimitGetEntities)
	q = buildUserWhereEqual(ormUser, q)
	err := q.Select(&users)

	return users, err
}
