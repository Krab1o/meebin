package auth

import (
	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/gin-gonic/gin"
)

func (h *handler) Logout(c *gin.Context) error {
	ctx := c.Request.Context()
	sessionId, ok := c.Get(shared.SessionIDJsonName)
	if !ok {
		return api.NewInternalError(nil, nil)
	}

	err := h.authService.Logout(ctx, sessionId.(uint64))
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}

	return nil
}
