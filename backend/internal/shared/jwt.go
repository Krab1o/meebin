package shared

import (
	"fmt"

	"github.com/Krab1o/meebin/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

const (
	UserIDJsonName    = "userId"
	SessionIDJsonName = "sessionId"
	RolesJsonName     = "roles"
)

type CustomAccessFields struct {
	UserID    uint64       `json:"userId"`
	SessionID uint64       `json:"sessionId"`
	Roles     []model.Role `json:"roles"`
}

type AccessClaims struct {
	CustomAccessFields `json:",inline"`
	jwt.RegisteredClaims
}

type CustomRefreshFields struct {
	SessionID uint64 `json:"sessionId"`
}

type RefreshClaims struct {
	CustomRefreshFields `json:",inline"`
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
