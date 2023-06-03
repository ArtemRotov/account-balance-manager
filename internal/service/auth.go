package service

import (
	"context"
	"errors"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/repository"
	"github.com/ArtemRotov/account-balance-manager/internal/repository/repoerrors"
	log "github.com/sirupsen/logrus"
)

type AuthService struct {
	userRepo repository.UserRepository
	hasher   PasswordHasher
}

func NewAuthService(r repository.UserRepository, h PasswordHasher) *AuthService {
	return &AuthService{
		userRepo: r,
		hasher:   h,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, username, password string) (int, error) {
	u := &model.User{
		Username: username,
		Password: s.hasher.Hash(password),
	}

	id, err := s.userRepo.CreateUser(ctx, u)
	if err != nil {
		if errors.Is(err, repoerrors.ErrAlreadyExists) {
			return 0, ErrUserAlreadyExists
		}
		log.Errorf("AuthService.CreateUser - cannot create user %v", err)
		return 0, err
	}

	return id, nil
}
