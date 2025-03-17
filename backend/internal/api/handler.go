package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler func(c *gin.Context) error

func logError(err error) {
	var logger strings.Builder
	logger.WriteString(fmt.Sprintf("%s\n", time.Now().Format(time.DateTime)))
	for err != nil {
		logger.WriteString(fmt.Sprintf("  %s", err.Error()))
		err = errors.Unwrap(err)
		if err != nil {
			logger.WriteString("\n")
		}
	}
	fmt.Println(logger.String())
}

func MakeHandler(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			c.Abort()
			var apiError *Error
			if errors.As(err, &apiError) {
				logError(apiError)
				c.JSON(apiError.StatusCode, apiError)
				return
			}
			apiError.Message = "Unknown Error"
			c.JSON(http.StatusInternalServerError, apiError)
			return
		}

	}
}
