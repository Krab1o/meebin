package auth

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/api/auth/converter"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/gin-gonic/gin"
)

func (h *handler) Profile(c *gin.Context) error {
	ctx := c.Request.Context()
	userId, ok := c.Get(shared.UserIDJsonName)
	if !ok {
		return api.NewInternalError(nil, "Unable to parse userId")
	}
	user, err := h.authService.Profile(ctx, userId.(uint64))
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}
	dtoUser := converter.UserServiceToDTO(user)
	c.JSON(http.StatusOK, dtoUser)
	return nil
}
