package storage

import (
	"github.com/voltento/users-info/app/model"
)

type (
	UsersFunc func(*model.User) ([]model.User, error)
	UserFunc  func(userId string) (*model.User, error)
	DropUser  func(userId string) error
)

type StorageMock struct {
	usersFunc UsersFunc
	userFunc  UserFunc
	dropUser  DropUser
}

func (s *StorageMock) DropUser(userId string) error {
	return s.dropUser(userId)
}

func NewStorageMock(usersFunc UsersFunc, userFunc UserFunc, dropUser DropUser) Storage {
	return &StorageMock{usersFunc: usersFunc, userFunc: userFunc, dropUser: dropUser}
}

func (s *StorageMock) Users(user *model.User) ([]model.User, error) {
	return s.usersFunc(user)
}

func (s *StorageMock) User(userId string) (*model.User, error) {
	return s.userFunc(userId)
}

func (s *StorageMock) Stop() error {
	return nil
}
