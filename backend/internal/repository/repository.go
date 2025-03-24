package repository

import (
	"context"
	"fmt"

	"github.com/Krab1o/meebin/internal/model"
	rmodelEvent "github.com/Krab1o/meebin/internal/model/event/r_model"
	rmodelUser "github.com/Krab1o/meebin/internal/model/user/r_model"
)

func Col(table, column string) string {
	return fmt.Sprintf("%s.%s", table, column)
}

type UserRepository interface {
	AddUser(ctx context.Context, user *rmodelUser.User, roleId uint64) (uint64, error)
	GetCredsByEmail(ctx context.Context, email string) (*rmodelUser.User, error)
	GetById(ctx context.Context, userId uint64) (*rmodelUser.User, error)
	DeleteById(ctx context.Context, userId uint64) error
	List(ctx context.Context) ([]rmodelUser.User, error)
	UpdateStats(ctx context.Context, userId uint64, stats *rmodelUser.Stats) error
	UpdateCreds(ctx context.Context, userId uint64, creds *rmodelUser.Creds) error
	UpdatePersonalData(
		ctx context.Context,
		userId uint64,
		data *rmodelUser.PersonalData,
	) error
}
type SessionRepository interface {
	AddSession(ctx context.Context, session *rmodelUser.Session) (uint64, error)
	DeleteSession(ctx context.Context, sessionId uint64) error
	FindSession(ctx context.Context, sessionId uint64) (*rmodelUser.Session, error)
}

type RoleRepository interface {
	GetRolesByTitle(ctx context.Context, roles []model.Role) (uint64, error)
	GetUserRolesById(ctx context.Context, userId uint64) ([]model.Role, error)
}

// type UserDataRepository interface {
// }
// type UserStatsRepository interface {
// }

// TODO: think of possible renamings in model's folder
type EventRepository interface {
	Add(ctx context.Context, newEvent *rmodelEvent.Event) (uint64, error)
	GetById(ctx context.Context, eventId uint64) (*rmodelEvent.Event, error)
	List(ctx context.Context) ([]rmodelEvent.Event, error)
	UpdateEvent(ctx context.Context, eventId uint64, event *rmodelEvent.Event) error
	UpdateEventData(ctx context.Context, eventId uint64, eventData *rmodelEvent.EventData) error
	// Update(ctx context.Context, id uint64) (*rmodelEvent.Event, error)
	Delete(ctx context.Context, id uint64) error
}
