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
			return api.NewUnauthorizedError("No authorization token", nil)
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return api.NewUnauthorizedError(
				"Authorization token should be formatted as \"Bearer <token>\"",
				nil,
			)
		}

		tokenString := parts[1]

		claims := &shared.AccessClaims{}
		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			shared.ParseFunction(jwtSecret),
		)
		//TODO: check different types of errors, add messages
		if err != nil {
			return api.NewUnauthorizedError("Unexpected singning method", err)
		}
		if !token.Valid {
			return api.NewUnauthorizedError("Invalid token", nil)
		}
		c.Set(shared.UserIDJsonName, claims.UserID)
		c.Set(shared.SessionIDJsonName, claims.SessionID)
		return nil
	}
}
