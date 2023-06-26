package pgdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/repository/repoerrors"
)

type ReservationRepo struct {
	db *sql.DB
}

func NewReservationRepo(db *sql.DB) *ReservationRepo {
	return &ReservationRepo{
		db: db,
	}
}

func (r *ReservationRepo) CreateReservation(ctx context.Context, rsv *model.Reservation) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, repoerrors.ErrCannotStartTransaction
	}

	rows := tx.QueryRow(
		`SELECT * FROM reservations WHERE 
		account_id = $1 AND  service_id = $2 AND order_id = $3 AND amount = $4`,
		rsv.AccountId, rsv.ServiceId, rsv.OrderId, rsv.Amount)
	if rows.Err() == nil {
		tx.Rollback()
		return 0, repoerrors.ErrAlreadyExists
	} else if !errors.Is(rows.Err(), sql.ErrNoRows) {
		tx.Rollback()
		return 0, rows.Err()
	}

	id := 0
	if err := tx.QueryRow(
		`INSERT INTO (account_id, service_id, order_id, amount) VALUES ($1, $2, $3, $4) RETURNING id`,
		rsv.AccountId, rsv.ServiceId, rsv.OrderId, rsv.Amount,
	).Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return id, nil
}
