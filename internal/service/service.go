package service

import (
	"golang.org/x/exp/slog"
	"time"

	"github.com/ArtemRotov/account-balance-manager/internal/repository"
)

type Services struct {
	Auth        *AuthService
	Account     *AccountService
	Reservation *ReservationService
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
