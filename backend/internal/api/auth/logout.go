package auth

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/gin-gonic/gin"
)

// @Tags			Auth
// @Summary		Logout
// @Schemes		http
// @Description	Disables refresh token so you won't able to use it in /auth/refresh method
// @Accept			json
// @Produce		json
// @Security		jwtToken
// @Success		200
// @Failure		400	{object}	api.Error
// @Failure		401	{object}	api.Error
// @Failure		500	{object}	api.Error
// @Router			/auth/logout [post]
func (h *Handler) Logout(c *gin.Context) error {
	ctx := c.Request.Context()
	sessionId, ok := c.Get(shared.SessionIDJsonName)
	if !ok {
		return api.NewInternalError(nil, "Unable to parse sessionId")
	}
	err := h.authService.Logout(ctx, sessionId.(uint64))
	if err != nil {
		switch {
		default:
			return api.NewInternalError(err)
		}
	}
	c.Status(http.StatusOK)
	return nil
}
