package storage

import "github.com/voltento/users-info/app/model"

type Storage interface {
	Users(user *model.User) ([]model.User, error)
	User(userId string) (*model.User, error)
	Stop() error
}
