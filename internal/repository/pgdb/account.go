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
