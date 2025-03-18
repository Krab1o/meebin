package user

import (
	"errors"
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/gin-gonic/gin"
)

// @Tags			User
// @Summary		Deletes user
// @Schemes		http
// @Description	Deletes user by ID
// @Accept			json
// @Produce		json
// @Security		jwtToken
// @Param			user_id	path	int	true	"User ID"
// @Success		204
// @Failure		400	{object}	api.Error
// @Failure		401	{object}	api.Error
// @Failure		404	{object}	api.Error
// @Failure		500	{object}	api.Error
// @Router			/users/{user_id} [delete]
func (h *handler) DeleteUser(c *gin.Context) error {
	ctx := c.Request.Context()
	id, ok := c.Get(shared.UserIDJsonName)
	if !ok {
		return api.NewInternalError(nil)
	}
	err := h.userService.Delete(ctx, id.(uint64))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			return api.NewNotFoundError(err, "Not found user to delete")
		default:
			return api.NewInternalError(err)
		}
	}
	c.Status(http.StatusNoContent)
	return nil
}
