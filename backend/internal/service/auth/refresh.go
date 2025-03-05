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
	//TODO: parse refresh token

	//TODO: move to one place (refactor)
	claims := &shared.RefreshClaims{}
	token, err := jwt.ParseWithClaims(
		refreshToken,
		claims,
		shared.ParseFunction(s.jwtConf.Secret()),
	)
	if err != nil {
		return "", service.NewUnauthorizedError(err)
	}
	if !token.Valid {
		return "", service.NewUnauthorizedError(nil)
	}
	repoSession, err := s.sessionRepo.FindSession(ctx, nil, claims.SessionID)
	if err != nil {
		return "", parseDBError(err)
	}
	timeNow := time.Now()
	accessToken, err := helper.GenerateAccessToken(
		repoSession.UserId,
		repoSession.SessionId,
		timeNow,
		s.jwtConf.Secret(),
		s.jwtConf.AccessTimeout(),
	)
	if err != nil {
		return "", service.NewInternalError(err)
	}
	return accessToken, nil
}
