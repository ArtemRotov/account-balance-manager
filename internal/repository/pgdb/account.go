package pgdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/repository/repoerrors"
)

type AccountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{
		db: db,
	}
}

func (r *AccountRepo) AccountByUserId(ctx context.Context, userId int) (*model.Account, error) {
	a := &model.Account{}

	if err := r.db.QueryRow(
		"SELECT id, user_id, balance, created_at FROM accounts WHERE user_id = $1",
		userId,
	).Scan(&a.Id, &a.UserId, &a.Balance, &a.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repoerrors.ErrNotFound
		}
		return nil, err
	}

	return a, nil
}

func (r *AccountRepo) DepositByUserId(ctx context.Context, userId, amount int) (*model.Account, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return nil, err
	}

	a := &model.Account{}
	if err := tx.QueryRow(
		"SELECT id, user_id, balance, created_at FROM accounts where user_id = $1",
		userId,
	).Scan(&a.Id, &a.UserId, &a.Balance, &a.CreatedAt); err != nil {
		tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repoerrors.ErrNotFound
		}
		return nil, err
	}

	a.Balance += amount
	if row := tx.QueryRow(
		"UPDATE accounts SET balance = $1 WHERE id = $2",
		a.Balance,
		a.Id,
	); row.Err() != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return a, nil
}
