package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler func(c *gin.Context) error

func MakeHandler(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			var apiError *Error
			if errors.As(err, &apiError) {
				c.JSON(apiError.StatusCode, apiError)
				return
			}
			apiError.Message = "Unknown Error"
			c.JSON(http.StatusInternalServerError, apiError)
			return
		}

	}
}
