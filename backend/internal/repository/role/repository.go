package role

import (
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.RoleRepository {
	return &repo{db: db}
}
