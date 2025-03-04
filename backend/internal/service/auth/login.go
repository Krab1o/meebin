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
	accessToken, err := helper.GenerateAccessToken(
		userId,
		timeNow,
		s.jwtConf.Secret(),
		s.jwtConf.AccessTimeout(),
	)
	if err != nil {
		return nil, service.NewInternalError(err)
	}
	refreshToken, expirationTime, err := helper.GenerateRefreshToken(
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
	s.sessionRepo.AddSession(ctx, nil, repoSession)
	return &smodel.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
