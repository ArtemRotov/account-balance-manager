package pgdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ArtemRotov/account-balance-manager/internal/model"
	"github.com/ArtemRotov/account-balance-manager/internal/repository/repoerrors"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

//	Создается не только User, но и аккаунт связаный с ним. Наверное это плохо
func (r *UserRepo) CreateUser(ctx context.Context, user *model.User) (int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	id := 0
	err = tx.QueryRow(
		"SELECT id FROM users WHERE username = $1",
		user.Username,
	).Scan(&id)

	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return 0, err
	}

	if id != 0 {
		tx.Rollback()
		return 0, repoerrors.ErrAlreadyExists
	}

	if err := tx.QueryRow(
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id",
		user.Username,
		user.Password,
	).Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	if row := tx.QueryRow(
		"INSERT INTO accounts (user_id) VALUES ($1) RETURNING id",
		id,
	); row.Err() != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()
	return id, nil
}

func (r *UserRepo) UserByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error) {
	u := &model.User{}

	if err := r.db.QueryRow(
		"SELECT id, username, password, created_at FROM users where username = $1 AND password = $2",
		username,
		password,
	).Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repoerrors.ErrNotFound
		}
		return nil, err //something wrong
	}

	return u, nil
}
