package connectors

import "github.com/voltento/users-info/app/model"

type Storage interface {
	Users() (error, []model.User)
	Stop() error
}
