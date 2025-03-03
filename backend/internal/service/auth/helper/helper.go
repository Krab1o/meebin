package helper

import (
	"time"

	"github.com/Krab1o/meebin/internal/shared"
	"github.com/golang-jwt/jwt/v5"
)

//TODO: add roles to token

func GenerateAccessToken(userID uint64, timeNow time.Time, jwtSecret []byte, jwtTimeout int) (string, error) {
	timeNow := time.Now()
	expirationTime := timeNow.Add(time.Duration(jwtTimeout) * time.Minute)
	claims := shared.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(timeNow),
			// Subject:   "user_auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(userID uint64, timeNow time.Time, jwtSecret []byte, jwtTimeout int) (string, error) {
	timeNow := time.Now()
	expirationTime := timeNow.Add(time.Duration(jwtTimeout) * time.Hour)
	claims := shared.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(timeNow),
			// Subject:   "user_auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
