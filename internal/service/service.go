package service

import (
	"context"

	"github.com/ArtemRotov/account-balance-manager/internal/repository"
)

type Services struct {
	Auth
}

func NewServices(deps *ServiceDeps) *Services {
	return &Services{
		Auth: NewAuthService(deps.repo.UserRepository, deps.hasher),
	}
}

type ServiceDeps struct {
	repo   *repository.Repositories
	hasher PasswordHasher
}

func NewServicesDeps(repo *repository.Repositories, h PasswordHasher) *ServiceDeps {
	return &ServiceDeps{
		repo:   repo,
		hasher: h,
	}
}

type PasswordHasher interface {
	Hash(password string) string
}

type Auth interface {
	CreateUser(ctx context.Context, username, password string) (int, error)
}
