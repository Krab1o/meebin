package auth

import (
	"context"
	"log"
	"time"

	"github.com/Krab1o/meebin/internal/service"
	"github.com/Krab1o/meebin/internal/service/auth/helper"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/golang-jwt/jwt/v5"
)

// TODO: check if session with current refresh token exists
// if yes -> generate access token, return
// if no  -> error

func (s *serv) Refresh(ctx context.Context, refreshToken string) (string, error) {
	//TODO: move to one place (refactor)
	claims := &shared.RefreshClaims{}
	token, err := jwt.ParseWithClaims(
		refreshToken,
		claims,
		shared.ParseFunction(s.jwtConfig.Secret()),
	)
	if err != nil {
		return "", service.NewUnauthorizedError(err)
	}
	if !token.Valid {
		return "", service.NewUnauthorizedError(nil)
	}
	repoSession, err := s.sessionRepository.FindById(ctx, claims.SessionID)
	if err != nil {
		return "", service.ErrorDBToService(err)
	}
	roles, err := s.roleRepository.ListUserRolesById(ctx, repoSession.UserId)
	if err != nil {
		return "", service.ErrorDBToService(err)
	}
	log.Println(claims.CustomRefreshFields)
	log.Println(roles)
	timeNow := time.Now()
	accessToken, err := helper.GenerateAccessToken(
		shared.CustomAccessFields{
			UserID:    repoSession.UserId,
			SessionID: repoSession.SessionId,
			Roles:     roles,
		},
		timeNow,
		s.jwtConfig.Secret(),
		s.jwtConfig.AccessTimeout(),
	)
	if err != nil {
		return "", service.NewInternalError(err)
	}
	return accessToken, nil
}
