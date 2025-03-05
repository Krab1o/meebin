package repository

import (
	"context"

	rmodel "github.com/Krab1o/meebin/internal/struct/r_model"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	AddUser(context.Context, pgx.Tx, *rmodel.User) (uint64, error)
	FindUser(context.Context, pgx.Tx, *rmodel.Creds) (uint64, error)
	// GetById(id int64) (*user.User, error)
	// List() ([]user.User, error)
	// Update(id int64) error
	// Delete(id int64) error
}
type UserSessionRepository interface {
	AddSession(context.Context, pgx.Tx, *rmodel.Session) (uint64, error)
	DeleteSession(context.Context, pgx.Tx, uint64) error
	FindSession(context.Context, pgx.Tx, uint64) (*rmodel.Session, error)
}

// type UserDataRepository interface {
// }
// type UserStatsRepository interface {
// }

type EventRepository interface {
	// Add(*event.Event) (int64, error)
	// GetById(id int64) (*event.Event, error)
	// List() ([]event.Event, error)
	// Update(id int64) error
	// Delete(id int64) error
}
