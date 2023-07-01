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

type ReservationService struct {
	reservationRepo repository.ReservationRepository
	log             *slog.Logger
}

func NewReservationService(r repository.ReservationRepository, log *slog.Logger) *ReservationService {
	return &ReservationService{
		reservationRepo: r,
		log:             log,
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
			s.log.Error(fmt.Sprintf("ReservationService.CreateReservation - reservation exists %v", err))
			return nil, ErrReservationAlreadyExists
		} else if errors.Is(err, repoerrors.ErrInsufficientBalance) {
			s.log.Error(fmt.Sprintf("ReservationService.CreateReservation - not enough money %v", err))
			return nil, ErrNotEnoughMoney
		}
		s.log.Error(fmt.Sprintf("ReservationService.CreateReservation - repoerror %v", err))
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
			s.log.Error(fmt.Sprintf("ReservationService.Revenue - reservation not found %v", err))
			return ErrReservationNotFound
		}
		s.log.Error(fmt.Sprintf("ReservationService.Revenue - repoerror %v", err))
		return err
	}

	return nil
}

func (s *ReservationService) Refund(ctx context.Context, account_id, service_id, order_id, amount int) error {
	r := &model.Reservation{
		AccountId: account_id,
		ServiceId: service_id,
		OrderId:   order_id,
		Amount:    amount,
	}

	err := s.reservationRepo.Refund(ctx, r)
	if err != nil {
		if errors.Is(err, repoerrors.ErrNotFound) {
			s.log.Error(fmt.Sprintf("ReservationService.Refund - reservation not found %v", err))
			return ErrReservationNotFound
		}
		s.log.Error(fmt.Sprintf("ReservationService.Refund - repoerror %v", err))
		return err
	}

	return nil
}
