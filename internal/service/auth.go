package service

import (
	"context"

	"github.com/ArtemRotov/account-balance-manager/internal/repository"
)

type AuthService struct {
	repo   repository.UserRepository
	hasher PasswordHasher
}

func NewAuthService(r repository.UserRepository, h PasswordHasher) *AuthService {
	return &AuthService{
		repo:   r,
		hasher: h,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, username, password string) (int, error) {

	return 0, nil
}
