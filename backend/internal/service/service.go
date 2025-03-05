package service

import (
	"context"

	smodel "github.com/Krab1o/meebin/internal/struct/s_model"
)

type EventService interface {
}

type UserService interface {
}

type AuthService interface {
	Login(ctx context.Context, user *smodel.Creds) (*smodel.Tokens, error)
	Logout(ctx context.Context, sessionId uint64) error
	// Profile(ctx context.Context) error
	Refresh(ctx context.Context, refreshToken string) (string, error)
	Register(ctx context.Context, user *smodel.User) (*smodel.Tokens, error)
}
