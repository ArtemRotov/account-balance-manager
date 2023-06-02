package repository

import (
	"context"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (int, error)
}
