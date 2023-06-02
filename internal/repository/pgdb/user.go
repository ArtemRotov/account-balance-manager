package pgdb

import (
	"context"
	"database/sql"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *model.User) (int, error) {
	return 0, nil
}
