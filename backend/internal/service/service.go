package service

import (
	"context"

	smodel "github.com/Krab1o/meebin/internal/model/s_model"
)

type AuthService interface {
	Login(ctx context.Context, user *smodel.Creds) (*smodel.Tokens, error)
	Logout(ctx context.Context, sessionId uint64) error
	Refresh(ctx context.Context, refreshToken string) (string, error)
	Register(ctx context.Context, user *smodel.User) (*smodel.Tokens, error)
}

type UserService interface {
	ListUser(ctx context.Context) ([]smodel.User, error)
	GetUser(ctx context.Context, userId uint64) (*smodel.User, error)
	Delete(ctx context.Context, userId uint64) error
	Update(ctx context.Context, user *smodel.User) (*smodel.User, error)
}

type EventService interface {
}
