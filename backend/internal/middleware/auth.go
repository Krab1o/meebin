package middleware

import (
	"strings"

	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(jwtSecret []byte) api.Handler {
	return func(c *gin.Context) error {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			return api.NewUnauthorizedError(nil, "No authorization token")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return api.NewUnauthorizedError(
				nil,
				"Authorization token should be formatted as \"Bearer <token>\"",
			)
		}

		tokenString := parts[1]

		claims := &shared.AccessClaims{}
		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			shared.ParseFunction(jwtSecret),
		)
		if err != nil {
			return api.NewUnauthorizedError(err, "Unexpected singning method")
		}
		if !token.Valid {
			return api.NewUnauthorizedError(nil, "Invalid token")
		}
		c.Set(shared.UserIDJsonName, claims.UserID)
		c.Set(shared.SessionIDJsonName, claims.SessionID)
		c.Set(shared.RolesJsonName, claims.Roles)
		return nil
	}
}
