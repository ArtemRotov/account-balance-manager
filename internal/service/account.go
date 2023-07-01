package service

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/exp/slog"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/repository"
	"github.com/ArtemRotov/account-balance-manager/internal/repository/repoerrors"
)

type AccountService struct {
	accountRepo repository.AccountRepository
	log         *slog.Logger
}

func NewAccountService(accountRepo repository.AccountRepository, log *slog.Logger) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
		log:         log,
	}
}

func (s *AccountService) AccountByUserId(ctx context.Context, userId int) (*model.Account, error) {
	a, err := s.accountRepo.AccountByUserId(ctx, userId)
	if err != nil {
		if errors.Is(err, repoerrors.ErrNotFound) {
			s.log.Error("AccountService.AccountByUserId - account not found %v", err)
			return nil, ErrAccountNotFound
		}
		s.log.Error("AccountService.AccountByUserId - repoerror %v", err)
		return nil, err
	}

	return a, nil
}

func (s *AccountService) DepositByUserId(ctx context.Context, userId, amount int) (*model.Account, error) {
	a, err := s.accountRepo.DepositByUserId(ctx, userId, amount)
	if err != nil {
		if errors.Is(err, repoerrors.ErrNotFound) {
			s.log.Error(fmt.Sprintf("AccountService.Deposit - account not found %v", err))
			return nil, ErrAccountNotFound
		}
		s.log.Error(fmt.Sprintf("AccountService.Deposit - repoerror %v", err))
		return nil, err
	}

	return a, nil
}
