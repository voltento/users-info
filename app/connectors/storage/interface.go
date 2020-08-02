package storage

import (
	"context"
	"github.com/voltento/users-info/app/model"
)

type Storage interface {
	Users(user *model.User) ([]model.User, error)
	User(userId string) (*model.User, error)
	DropUser(userId string) error
	UpdateUser(user *model.User) error
	AddUser(user *model.User) error
	Stop() error
	Check(ctx context.Context) error
}
