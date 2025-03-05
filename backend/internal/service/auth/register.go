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
		return nil, service.ErrorDBToService(err, nil)
	}
	//TODO: add transaction
	userId, err := s.userRepo.AddUser(ctx, nil, repoUser)
	if err != nil {
		return nil, service.ErrorDBToService(err, nil)
	}

	// Creating database row with session
	timeNow := time.Now()
	refreshExpirationTime := timeNow.Add(time.Duration(s.jwtConf.RefreshTimeout()) * time.Hour)
	repoSession := &rmodel.Session{
		UserId:         userId,
		ExpirationTime: refreshExpirationTime,
	}
	sessionId, err := s.sessionRepo.AddSession(ctx, nil, repoSession)
	if err != nil {
		return nil, service.ErrorDBToService(err, nil)
	}

	// Write database data to refresh token
	refreshToken, err := authHelper.GenerateRefreshToken(
		sessionId,
		refreshExpirationTime,
		timeNow,
		s.jwtConf.Secret(),
	)
	if err != nil {
		return nil, service.NewInternalError(err)
	}

	// Generate access token
	accessToken, err := authHelper.GenerateAccessToken(
		userId,
		sessionId,
		timeNow,
		s.jwtConf.Secret(),
		s.jwtConf.AccessTimeout(),
	)
	if err != nil {
		return nil, service.NewInternalError(err)
	}

	return &smodel.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
