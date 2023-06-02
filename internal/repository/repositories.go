package repository

import (
	"database/sql"

	"github.com/ArtemRotov/account-balance-manager/internal/repository/pgdb"
)

type Repositories struct {
	UserRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository: pgdb.NewUserRepo(db),
	}
}
