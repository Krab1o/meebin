package service

import (
	"context"

	smodel "github.com/Krab1o/meebin/internal/model/s_model"
)

type EventService interface {
}

type UserService interface {
}

type AuthService interface {
	Login(ctx context.Context, id uint64) (*smodel.Tokens, error)
	// Logout(ctx context.Context) error
	// Profile(ctx context.Context) error
	Refresh(ctx context.Context, id uint64) (*smodel.AccessToken, error)
	Register(ctx context.Context, userCreds smodel.User) error
}
