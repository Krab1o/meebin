package auth

import (
	"context"
	"time"

	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/Krab1o/meebin/internal/service/auth/helper"
	"github.com/Krab1o/meebin/internal/shared"
)

// TODO: remove needing username
func (s *serv) Login(ctx context.Context, creds *smodel.Creds) (*smodel.Tokens, error) {
	repoUser, err := s.userRepository.GetCredsByEmail(ctx, creds.Email)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}
	ok := helper.VerifyPassword(
		repoUser.Creds.HashedPassword,
		creds.Password,
	)
	if !ok {
		return nil, service.NewUnauthorizedError(err)
	}

	timeNow := time.Now()
	refreshExpirationTime := timeNow.Add(time.Duration(s.jwtConfig.RefreshTimeout()) * time.Hour)
	repoSession := &rmodel.Session{
		UserId:         repoUser.Id,
		ExpirationTime: refreshExpirationTime,
	}
	sessionId, err := s.sessionRepository.AddSession(ctx, repoSession)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}

	refreshToken, err := helper.GenerateRefreshToken(
		shared.CustomRefreshFields{
			SessionID: sessionId,
		},
		refreshExpirationTime,
		time.Now(),
		s.jwtConfig.Secret(),
	)
	if err != nil {
		return nil, service.NewInternalError(err)
	}

	roles, err := s.roleRepository.GetUserRolesById(ctx, repoUser.Id)
	if err != nil {
		return nil, service.ErrorDBToService(err)
	}

	accessToken, err := helper.GenerateAccessToken(
		shared.CustomAccessFields{
			UserID:    repoUser.Id,
			SessionID: sessionId,
			Roles:     roles,
		},
		timeNow,
		s.jwtConfig.Secret(),
		s.jwtConfig.AccessTimeout(),
	)
	if err != nil {
		return nil, service.NewInternalError(err)
	}
	return &smodel.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
