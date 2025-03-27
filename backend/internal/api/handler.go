package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler func(c *gin.Context) error

func MakeHandler(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			fmt.Println(err)
			c.Abort()
			var apiError *Error
			//TODO: add validation type error
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
