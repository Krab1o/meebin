package auth

import "github.com/gin-gonic/gin"

func (h *handler) Logout(c *gin.Context) error {
	_ = c.Request.Context()
	return nil
}
