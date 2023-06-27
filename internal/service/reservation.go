package service

import (
	"context"
	"errors"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/repository"
	"github.com/ArtemRotov/account-balance-manager/internal/repository/repoerrors"
	"github.com/sirupsen/logrus"
)

type ReservationService struct {
	reservationRepo repository.ReservationRepository
}

func NewReservationService(r repository.ReservationRepository) *ReservationService {
	return &ReservationService{
		reservationRepo: r,
	}
}

func (s *ReservationService) CreateReservation(ctx context.Context, account_id,
	service_id, order_id, amount int) (*model.Reservation, error) {

	r := &model.Reservation{
		AccountId: account_id,
		ServiceId: service_id,
		OrderId:   order_id,
		Amount:    amount,
	}

	id, err := s.reservationRepo.CreateReservation(ctx, r)
	if err != nil {
		if errors.Is(err, repoerrors.ErrAlreadyExists) {
			logrus.Errorf("ReservationService.CreateReservation - reservation exists %v", err)
			return nil, ErrReservationAlreadyExists
		} else if errors.Is(err, repoerrors.ErrInsufficientBalance) {
			logrus.Errorf("ReservationService.CreateReservation - not enough money %v", err)
			return nil, ErrNotEnoughMoney
		}
		logrus.Errorf("ReservationService.CreateReservation - repoerror %v", err)
		return nil, err
	}

	r.Id = id
	return r, nil
}

func (s *ReservationService) Revenue(ctx context.Context, account_id, service_id, order_id, amount int) error {
	r := &model.Reservation{
		AccountId: account_id,
		ServiceId: service_id,
		OrderId:   order_id,
		Amount:    amount,
	}

	err := s.reservationRepo.Revenue(ctx, r)
	if err != nil {
		if errors.Is(err, repoerrors.ErrNotFound) {
			logrus.Errorf("ReservationService.Revenue - reservation not found %v", err)
			return ErrReservationNotFound
		}
		logrus.Errorf("ReservationService.Revenue - repoerror %v", err)
		return err
	}

	return nil
}
