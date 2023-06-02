package service

import (
	"context"

	"github.com/ArtemRotov/account-balance-manager/internal/repository"
)

type Services struct {
	Auth
}

func NewServices(repo *repository.Repositories) *Services {
	return &Services{
		Auth: NewAuthService(repo.UserRepository),
	}
}

type Auth interface {
	CreateUser(ctx context.Context, username, password string) (int, error)
}
