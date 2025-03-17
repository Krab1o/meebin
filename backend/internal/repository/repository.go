package repository

import (
	"context"
	"fmt"

	"github.com/Krab1o/meebin/internal/model"
	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	"github.com/jackc/pgx/v5"
)

func Col(table, column string) string {
	return fmt.Sprintf("%s.%s", table, column)
}

type UserRepository interface {
	AddUser(ctx context.Context, tx pgx.Tx, user *rmodel.User, roleId uint64) (uint64, error)
	GetCredsByEmail(ctx context.Context, tx pgx.Tx, email string) (*rmodel.User, error)
	GetById(ctx context.Context, tx pgx.Tx, id uint64) (*rmodel.User, error)
	DeleteById(ctx context.Context, tx pgx.Tx, id uint64) error
	List(ctx context.Context, tx pgx.Tx) ([]rmodel.User, error)
	UpdateStats(ctx context.Context, tx pgx.Tx, userId uint64, stats *rmodel.Stats) error
	UpdateCreds(ctx context.Context, tx pgx.Tx, userId uint64, creds *rmodel.Creds) error
	UpdatePersonalData(
		ctx context.Context,
		tx pgx.Tx,
		userId uint64,
		data *rmodel.PersonalData,
	) error
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
