package middleware

import (
	"fmt"
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

		claims := &shared.Claims{}
		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method")
				}
				return jwtSecret, nil
			},
		)
		//TODO: check different types of errors, add messages
		if err != nil {
			return api.NewUnauthorizedError("Unexpected singning method", err)
		}
		if !token.Valid {
			return api.NewUnauthorizedError("Invalid token", nil)
		}
		c.Set(shared.UserIDJsonName, claims.UserID)
		return nil
	}
}
