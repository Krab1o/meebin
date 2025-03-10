package auth

import (
	"context"
	"log"
	"time"

	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/s_model"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/Krab1o/meebin/internal/service/auth/helper"
	"github.com/Krab1o/meebin/internal/shared"
)

// TODO: fix not checking password
// TODO: remove needing username
func (s *authService) Login(ctx context.Context, creds *smodel.Creds) (*smodel.Tokens, error) {
	repoUser, err := s.userRepo.GetUserCredsByEmail(ctx, nil, creds.Email)
	if err != nil {
		return nil, service.ErrorDBToService(err, nil)
	}
	ok := helper.VerifyPassword(
		repoUser.Creds.HashedPassword,
		creds.Password,
	)
	if !ok {
		return nil, service.NewUnauthorizedError(err, "Wrong password")
	}

	timeNow := time.Now()
	refreshExpirationTime := timeNow.Add(time.Duration(s.jwtConf.RefreshTimeout()) * time.Hour)
	repoSession := &rmodel.Session{
		UserId:         repoUser.Id,
		ExpirationTime: refreshExpirationTime,
	}
	sessionId, err := s.sessionRepo.AddSession(ctx, nil, repoSession)
	if err != nil {
		return nil, service.ErrorDBToService(err, nil)
	}

	refreshToken, err := helper.GenerateRefreshToken(
		shared.CustomRefreshFields{
			SessionID: sessionId,
		},
		refreshExpirationTime,
		time.Now(),
		s.jwtConf.Secret(),
	)
	if err != nil {
		return nil, service.NewInternalError(err)
	}

	roles, err := s.roleRepo.GetUserRolesById(ctx, nil, repoUser.Id)
	if err != nil {
		return nil, service.ErrorDBToService(err, nil)
	}
	log.Println(roles)
	//TODO: add roles
	accessToken, err := helper.GenerateAccessToken(
		shared.CustomAccessFields{
			UserID:    repoUser.Id,
			SessionID: sessionId,
			Roles:     roles,
		},
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
