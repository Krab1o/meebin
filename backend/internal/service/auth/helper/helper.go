package helper

import (
	"time"

	"github.com/Krab1o/meebin/internal/shared"
	"github.com/golang-jwt/jwt/v5"
)

//TODO: add roles to token

func GenerateAccessToken(
	userID uint64,
	sessionID uint64,
	timeNow time.Time,
	jwtSecret []byte,
	jwtTimeout int,
) (string, error) {
	expirationTime := timeNow.Add(time.Duration(jwtTimeout) * time.Minute)
	claims := shared.AccessClaims{
		UserID:    userID,
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(timeNow),
			// Subject:   "user_auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(
	sessionID uint64,
	expirationTime time.Time,
	timeNow time.Time,
	jwtSecret []byte,
) (string, error) {
	claims := shared.RefreshClaims{
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(timeNow),
			// Subject:   "user_auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
