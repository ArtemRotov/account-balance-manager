package service

import (
	"context"
	"errors"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/repository"
	"github.com/ArtemRotov/account-balance-manager/internal/repository/repoerrors"
	"github.com/sirupsen/logrus"
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
			logrus.Errorf("AccountService.AccountByUserId - account not found %v", err)
			return nil, ErrAccountNotFound
		}
		logrus.Errorf("AccountService.AccountByUserId - repoerror %v", err)
		return nil, err
	}

	return a, nil
}

func (s *AccountService) DepositByUserId(ctx context.Context, userId, amount int) (*model.Account, error) {
	a, err := s.accountRepo.DepositByUserId(ctx, userId, amount)
	if err != nil {
		if errors.Is(err, repoerrors.ErrNotFound) {
			logrus.Errorf("AccountService.Deposit - account not found %v", err)
			return nil, ErrAccountNotFound
		}
		logrus.Errorf("AccountService.Deposit - repoerror %v", err)
		return nil, err
	}

	return a, nil
}
