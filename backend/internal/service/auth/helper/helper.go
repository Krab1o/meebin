package helper

import (
	"time"

	"github.com/Krab1o/meebin/internal/shared"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(hashedPassword string, candidatePassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
	return err == nil
}

func GenerateAccessToken(
	accessFields shared.CustomAccessFields,
	timeNow time.Time,
	jwtSecret []byte,
	jwtTimeout int,
) (string, error) {
	expirationTime := timeNow.Add(time.Duration(jwtTimeout) * time.Minute)
	claims := shared.AccessClaims{
		CustomAccessFields: accessFields,
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
	refreshFields shared.CustomRefreshFields,
	expirationTime time.Time,
	timeNow time.Time,
	jwtSecret []byte,
) (string, error) {
	claims := shared.RefreshClaims{
		CustomRefreshFields: refreshFields,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(timeNow),
			// Subject:   "user_auth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
