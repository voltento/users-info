package connectors

import "github.com/voltento/users-info/app/model"

type Storage interface {
	Users() ([]model.User, error)
	User(userId string) (*model.User, error)
	Stop() error
}
