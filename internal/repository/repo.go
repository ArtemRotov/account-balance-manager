package repository

import (
	"context"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (int, error)
	UserByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error)
}

type AccountRepository interface {
	AccountByUserId(ctx context.Context, userId int) (*model.Account, error)
	DepositByUserId(ctx context.Context, userId, amount int) (*model.Account, error)
}

type ReservationRepository interface {
	CreateReservation(ctx context.Context, rsv *model.Reservation) (int, error)
	Revenue(ctx context.Context, rsv *model.Reservation) error
}
