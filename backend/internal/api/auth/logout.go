package auth

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/gin-gonic/gin"
)

func (h *handler) Logout(c *gin.Context) error {
	ctx := c.Request.Context()
	sessionId, ok := c.Get(shared.SessionIDJsonName)
	if !ok {
		return api.NewInternalError(nil, "Unable to parse sessionId")
	}
	err := h.authService.Logout(ctx, sessionId.(uint64))
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}
	c.Status(http.StatusOK)
	return nil
}
