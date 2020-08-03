package consumer

import (
	"context"
	"github.com/voltento/users-info/app/model"
)

type Mock struct {
}

func NewMock() Consumer {
	return &Mock{}
}

func (m Mock) UserUpdated(*model.User) {

}

func (m Mock) Check(context.Context) error {
	return nil
}
