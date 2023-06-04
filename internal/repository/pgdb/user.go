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

func (r *UserRepo) CreateUser(ctx context.Context, user *model.User) (int, error) {
	id := 0
	err := r.db.QueryRow(
		"SELECT id FROM users WHERE username = $1",
		user.Username,
	).Scan(&id)

	if err != nil && err != sql.ErrNoRows { //something wrong
		return 0, err
	}

	if id != 0 {
		return 0, repoerrors.ErrAlreadyExists
	}

	if err := r.db.QueryRow(
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id",
		user.Username,
		user.Password,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) UserByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error) {
	u := &model.User{}

	err := r.db.QueryRow(
		"SELECT id, username, password, created_at FROM users where username = $1 AND password = $2",
		username,
		password,
	).Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repoerrors.ErrNotFound
		}
		return nil, err //something wrong
	}

	return u, nil
}
