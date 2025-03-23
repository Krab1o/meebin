package user

import (
	"github.com/Krab1o/meebin/internal/client/db"
	"github.com/Krab1o/meebin/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}
