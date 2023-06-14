package service

import (
	"context"
	"errors"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/repository"
	"github.com/ArtemRotov/account-balance-manager/internal/repository/repoerrors"
)

type AccountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}

func (s *AccountService) AccountByUserId(ctx context.Context, userId int) (*model.Account, error) {
	a, err := s.accountRepo.AccountByUserId(ctx, userId)
	if err != nil {
		if errors.Is(err, repoerrors.ErrNotFound) {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}

	return a, nil
}
