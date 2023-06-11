package service

import (
	"context"
	"time"

	"github.com/ArtemRotov/account-balance-manager/internal/repository"
)

type Services struct {
	Auth
}

func NewServices(deps *ServiceDeps) *Services {
	return &Services{
		Auth: NewAuthService(deps.repo.UserRepository, deps.hasher, deps.signKey, deps.tokenTTL),
	}
}

type ServiceDeps struct {
	repo     *repository.Repositories
	hasher   PasswordHasher
	signKey  string
	tokenTTL time.Duration
}

func NewServicesDeps(repo *repository.Repositories, h PasswordHasher, signKey string, tokenTTL time.Duration) *ServiceDeps {
	return &ServiceDeps{
		repo:     repo,
		hasher:   h,
		signKey:  signKey,
		tokenTTL: tokenTTL,
	}
}

type PasswordHasher interface {
	Hash(password string) string
}

type Auth interface {
	CreateUser(ctx context.Context, username, password string) (int, error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}
