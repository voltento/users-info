package storage

import (
	"context"
	"github.com/voltento/users-info/app/model"
)

type (
	UsersFunc  func(*model.User) ([]model.User, error)
	UserFunc   func(userId string) (*model.User, error)
	DropUser   func(userId string) error
	UpdateUser func(*model.User) error
	AddUser    func(*model.User) error
)

type StorageMock struct {
	usersFunc      UsersFunc
	userFunc       UserFunc
	dropUserFunc   DropUser
	updateUserFunc UpdateUser
	addUserFunc    AddUser
}

func (s *StorageMock) AddUser(user *model.User) error {
	return s.addUserFunc(user)
}

func (s *StorageMock) Check(_ context.Context) error {
	return nil
}

func (s *StorageMock) UpdateUser(user *model.User) error {
	return s.updateUserFunc(user)
}

func (s *StorageMock) DropUser(userId string) error {
	return s.dropUserFunc(userId)
}

func NewStorageMock(users UsersFunc, user UserFunc, drop DropUser, update UpdateUser, add AddUser) Storage {
	return &StorageMock{usersFunc: users, userFunc: user, dropUserFunc: drop, updateUserFunc: update, addUserFunc: add}
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
