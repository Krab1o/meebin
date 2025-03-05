package shared

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

const (
	UserIDJsonName    = "userId"
	SessionIDJsonName = "sessionId"
)

type AccessClaims struct {
	UserID    uint64 `json:"userId"`
	SessionID uint64 `json:"sessionId"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	SessionID uint64 `json:"sessionId"`
	jwt.RegisteredClaims
}

func ParseFunction(jwtSecret []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return jwtSecret, nil
	}
}
