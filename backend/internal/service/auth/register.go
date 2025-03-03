package auth

import (
	"context"
	"time"

	"github.com/Krab1o/meebin/internal/service/auth/converter"
	authHelper "github.com/Krab1o/meebin/internal/service/auth/helper"
	smodel "github.com/Krab1o/meebin/internal/struct/s_model"
)

const (
	startUtilizeCount = 0
	startReportCount  = 0
	startRating       = 0.0
)

// TODO: add error messages
func (s *authService) Register(ctx context.Context, user *smodel.User) (*smodel.Tokens, error) {
	user.Stats = &smodel.Stats{
		UtilizeCount: startUtilizeCount,
		ReportCount:  startReportCount,
		Rating:       startRating,
	}
	repoUser, err := converter.UserServiceToRepository(user)
	if err != nil {
		return nil, err
	}
	id, err := s.userRepo.Add(ctx, nil, repoUser)
	if err != nil {
		return nil, err
	}
	timeNow := time.Now()
	accessToken, err := authHelper.GenerateAccessToken(id, s.jwtConf.Secret(), s.jwtConf.AccessTimeout())
	if err != nil {
		return nil, err
	}
	refreshToken, err := authHelper.GenerateRefreshToken(id, s.jwtConf.Secret(), s.jwtConf.RefreshTimeout())
	if err != nil {
		return nil, err
	}
	//TODO: generate JWT-tokens
	return &smodel.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
