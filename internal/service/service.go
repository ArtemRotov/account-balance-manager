package service

import (
	"context"
	"golang.org/x/exp/slog"
	"time"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/repository"
)

type Services struct {
	Auth
	Account
	Reservation
}

func NewServices(deps *ServiceDeps) *Services {
	return &Services{
		Auth:        NewAuthService(deps.repo.UserRepository, deps.log, deps.hasher, deps.signKey, deps.tokenTTL),
		Account:     NewAccountService(deps.repo.AccountRepository, deps.log),
		Reservation: NewReservationService(deps.repo, deps.log),
	}
}

type ServiceDeps struct {
	repo     *repository.Repositories
	log      *slog.Logger
	hasher   PasswordHasher
	signKey  string
	tokenTTL time.Duration
}

func NewServicesDeps(repo *repository.Repositories, log *slog.Logger, h PasswordHasher, signKey string, tokenTTL time.Duration) *ServiceDeps {
	return &ServiceDeps{
		repo:     repo,
		log:      log,
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

type Account interface {
	AccountByUserId(ctx context.Context, userId int) (*model.Account, error)
	DepositByUserId(ctx context.Context, userId, amount int) (*model.Account, error)
}

type Reservation interface {
	CreateReservation(ctx context.Context, account_id, service_id, order_id, amount int) (*model.Reservation, error)
	Revenue(ctx context.Context, account_id, service_id, order_id, amount int) error
	Refund(ctx context.Context, account_id, service_id, order_id, amount int) error
}
