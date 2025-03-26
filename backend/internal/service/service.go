package service

import (
	"context"

	smodelEvent "github.com/Krab1o/meebin/internal/model/event/s_model"
	smodelUser "github.com/Krab1o/meebin/internal/model/user/s_model"
)

type AuthService interface {
	Login(ctx context.Context, user *smodelUser.Creds) (*smodelUser.Tokens, error)
	Logout(ctx context.Context, sessionId uint64) error
	Refresh(ctx context.Context, refreshToken string) (string, error)
	Register(ctx context.Context, user *smodelUser.User) (*smodelUser.Tokens, error)
}

type UserService interface {
	List(ctx context.Context) ([]smodelUser.User, error)
	Get(ctx context.Context, userId uint64) (*smodelUser.User, error)
	Delete(ctx context.Context, deleterId uint64, userId uint64) error
	Update(
		ctx context.Context,
		updaterId uint64,
		user *smodelUser.User,
	) (*smodelUser.User, error)
}

type EventService interface {
	List(ctx context.Context) ([]smodelEvent.Event, error)
	Get(ctx context.Context, eventId uint64) (*smodelEvent.Event, error)
	Delete(ctx context.Context, deleterId uint64, eventId uint64) error
	Update(
		ctx context.Context,
		updaterId uint64,
		event *smodelEvent.Event,
	) (*smodelEvent.Event, error)
	Create(ctx context.Context, callerId uint64, event *smodelEvent.Event) (uint64, error)
}
