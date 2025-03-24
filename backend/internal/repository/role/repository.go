package role

import (
	"github.com/Krab1o/meebin/internal/client/db"
	rep "github.com/Krab1o/meebin/internal/repository"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) rep.RoleRepository {
	return &repo{db: db}
}
