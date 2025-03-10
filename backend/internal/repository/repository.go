package repository

import (
	"context"

	"github.com/Krab1o/meebin/internal/model"
	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	AddUser(ctx context.Context, tx pgx.Tx, user *rmodel.User, roleId uint64) (uint64, error)
	GetUserCredsByEmail(ctx context.Context, tx pgx.Tx, email string) (*rmodel.User, error)
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

type RoleRepository interface {
	GetRolesByTitle(context.Context, pgx.Tx, []model.Role) (uint64, error)
	GetUserRolesById(ctx context.Context, tx pgx.Tx, userId uint64) ([]model.Role, error)
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
