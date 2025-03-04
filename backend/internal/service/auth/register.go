package auth

import (
	"context"
	"time"

	"github.com/Krab1o/meebin/internal/service"
	"github.com/Krab1o/meebin/internal/service/auth/converter"
	authHelper "github.com/Krab1o/meebin/internal/service/auth/helper"
	rmodel "github.com/Krab1o/meebin/internal/struct/r_model"
	smodel "github.com/Krab1o/meebin/internal/struct/s_model"
)

// TODO: add error messages
func (s *authService) Register(ctx context.Context, user *smodel.User) (*smodel.Tokens, error) {
	repoUser, err := converter.UserServiceToRepository(user)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}
	//TODO: add transaction
	userId, err := s.userRepo.AddUser(ctx, nil, repoUser)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}

	timeNow := time.Now()
	accessToken, err := authHelper.GenerateAccessToken(
		userId,
		timeNow,
		s.jwtConf.Secret(),
		s.jwtConf.AccessTimeout(),
	)
	if err != nil {
		return nil, service.NewInternalError(err)
	}
	refreshToken, expirationTime, err := authHelper.GenerateRefreshToken(
		userId,
		timeNow,
		s.jwtConf.Secret(),
		s.jwtConf.RefreshTimeout(),
	)
	if err != nil {
		return nil, service.NewInternalError(err)
	}

	repoSession := &rmodel.Session{
		UserId:         userId,
		RefreshToken:   refreshToken,
		ExpirationTime: expirationTime,
	}
	err = s.sessionRepo.AddSession(ctx, nil, repoSession)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}
	return &smodel.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
