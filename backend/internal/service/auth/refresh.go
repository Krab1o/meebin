package auth

import (
	"context"
	"time"

	"github.com/Krab1o/meebin/internal/service"
	"github.com/Krab1o/meebin/internal/service/auth/helper"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/golang-jwt/jwt/v5"
)

// TODO: check if session with current refresh token exists
// if yes -> generate access token, return
// if no  -> error

func (s *authService) Refresh(ctx context.Context, refreshToken string) (string, error) {
	//TODO: move to one place (refactor)
	claims := &shared.RefreshClaims{}
	token, err := jwt.ParseWithClaims(
		refreshToken,
		claims,
		shared.ParseFunction(s.jwtConf.Secret()),
	)
	if err != nil {
		return "", service.NewUnauthorizedError(err, "Unable to parse token")
	}
	if !token.Valid {
		return "", service.NewUnauthorizedError(nil, "Invalid token")
	}
	repoSession, err := s.sessionRepo.FindSession(ctx, nil, claims.SessionID)
	if err != nil {
		return "", service.ErrorDBToService(err, nil)
	}
	roles, err := s.roleRepo.GetUserRolesById(ctx, nil, repoSession.UserId)
	if err != nil {
		return "", service.ErrorDBToService(err, nil)
	}
	timeNow := time.Now()
	accessToken, err := helper.GenerateAccessToken(
		shared.CustomAccessFields{
			UserID:    repoSession.UserId,
			SessionID: repoSession.SessionId,
			Roles:     roles,
		},
		timeNow,
		s.jwtConf.Secret(),
		s.jwtConf.AccessTimeout(),
	)
	if err != nil {
		return "", service.NewInternalError(err)
	}
	return accessToken, nil
}
