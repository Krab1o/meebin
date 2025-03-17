package user

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteUser(c *gin.Context) error {
	ctx := c.Request.Context()
	id, ok := c.Get(shared.UserIDJsonName)
	if !ok {
		return api.NewInternalError(nil, "Service error")
	}
	err := h.userService.Delete(ctx, id.(uint64))
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}
	c.Status(http.StatusNoContent)
	return nil
}
