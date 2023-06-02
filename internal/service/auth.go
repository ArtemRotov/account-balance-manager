package service

import (
	"context"

	"github.com/ArtemRotov/account-balance-manager/internal/repository"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(r repository.UserRepository) *AuthService {
	return &AuthService{
		repo: r,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, username, password string) (int, error) {

	return 0, nil
}
