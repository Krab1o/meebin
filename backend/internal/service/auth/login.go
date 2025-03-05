package auth

import (
	"context"
	"time"

	"github.com/Krab1o/meebin/internal/service"
	"github.com/Krab1o/meebin/internal/service/auth/helper"
	rmodel "github.com/Krab1o/meebin/internal/struct/r_model"
	smodel "github.com/Krab1o/meebin/internal/struct/s_model"
)

func (s *authService) Login(ctx context.Context, creds *smodel.Creds) (*smodel.Tokens, error) {
	repoCreds := &rmodel.Creds{
		Username: creds.Username,
		Email:    creds.Email,
		Password: creds.Password,
	}
	userId, err := s.userRepo.FindUser(ctx, nil, repoCreds)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}

	timeNow := time.Now()
	refreshExpirationTime := timeNow.Add(time.Duration(s.jwtConf.RefreshTimeout()) * time.Hour)
	repoSession := &rmodel.Session{
		UserId:         userId,
		ExpirationTime: refreshExpirationTime,
	}
	sessionId, err := s.sessionRepo.AddSession(ctx, nil, repoSession)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}

	refreshToken, err := helper.GenerateRefreshToken(
		sessionId,
		refreshExpirationTime,
		time.Now(),
		s.jwtConf.Secret(),
	)
	if err != nil {
		return nil, service.NewInternalError(err)
	}

	accessToken, err := helper.GenerateAccessToken(
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
