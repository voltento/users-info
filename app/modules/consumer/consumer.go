package consumer

import (
	"context"
	"github.com/voltento/users-info/app/model"
)

// Interface for updates consumer
type Consumer interface {
	UserUpdated(user *model.User)
	Check(ctx context.Context) error
}
