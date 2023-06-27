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

	cnt := 0
	if err := tx.QueryRow(
		`SELECT COUNT(*) FROM reservations WHERE 
		account_id = $1 AND  service_id = $2 AND order_id = $3 AND amount = $4`,
		rsv.AccountId, rsv.ServiceId, rsv.OrderId, rsv.Amount,
	).Scan(&cnt); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			tx.Rollback()
			return 0, err
		}
	}
	if cnt > 0 {
		tx.Rollback()
		return 0, repoerrors.ErrAlreadyExists
	}

	balance := 0
	if err := tx.QueryRow(
		"SELECT balance FROM accounts WHERE user_id = $1",
		rsv.AccountId,
	).Scan(&balance); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			tx.Rollback()
			return 0, repoerrors.ErrNotFound
		}
		return 0, err
	}

	if balance-rsv.Amount < 0 {
		tx.Rollback()
		return 0, repoerrors.ErrInsufficientBalance
	}

	if _, err := tx.Exec(
		"UPDATE accounts SET balance = $1 WHERE user_id = $2",
		balance-rsv.Amount,
		rsv.AccountId,
	); err != nil {
		tx.Rollback()
		return 0, err
	}

	id := 0
	if err := tx.QueryRow(
		`INSERT INTO reservations (account_id, service_id, order_id, amount) VALUES ($1, $2, $3, $4) RETURNING id`,
		rsv.AccountId, rsv.ServiceId, rsv.OrderId, rsv.Amount,
	).Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return id, nil
}

func (r *ReservationRepo) Revenue(ctx context.Context, rsv *model.Reservation) error {
	tx, err := r.db.Begin()
	if err != nil {
		return repoerrors.ErrCannotStartTransaction
	}

	if err := tx.QueryRow(
		`SELECT id FROM reservations WHERE	account_id = $1 AND
											service_id = $2 AND
											order_id   = $3 AND
											amount 	   = $4`,
		rsv.AccountId, rsv.ServiceId, rsv.OrderId, rsv.Amount,
	).Scan(&rsv.Id); err != nil {
		tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			return repoerrors.ErrNotFound
		}
		return err
	}

	if _, err := tx.Exec(
		"DELETE FROM reservations WHERE id = $1",
		rsv.Id,
	); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *ReservationRepo) Refund(ctx context.Context, rsv *model.Reservation) error {
	tx, err := r.db.Begin()
	if err != nil {
		return repoerrors.ErrCannotStartTransaction
	}

	if err := tx.QueryRow(
		`SELECT id FROM reservations WHERE	account_id = $1 AND
											service_id = $2 AND
											order_id   = $3 AND
											amount 	   = $4`,
		rsv.AccountId, rsv.ServiceId, rsv.OrderId, rsv.Amount,
	).Scan(&rsv.Id); err != nil {
		tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			return repoerrors.ErrNotFound
		}
		return err
	}

	if _, err := tx.Exec(
		"UPDATE accounts SET balance = balance + $1 WHERE id = $2",
		rsv.Amount,
		rsv.AccountId,
	); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(
		"DELETE FROM reservations WHERE id = $1",
		rsv.Id,
	); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
