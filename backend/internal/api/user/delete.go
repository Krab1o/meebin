package user

import (
	"errors"
	"net/http"
	"strconv"

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
func (h *Handler) Delete(c *gin.Context) error {
	ctx := c.Request.Context()
	deleterId, ok := c.Get(shared.UserIDJsonName)
	if !ok {
		return api.NewInternalError(nil)
	}
	paramId, ok := c.Params.Get(api.ParamId)
	if !ok {
		return api.NewInternalError(nil, "Unable to parse query param")
	}
	userIdToDelete, err := strconv.Atoi(paramId)
	if err != nil {
		return api.NewInternalError(nil, "Unable to convert query param")
	}
	err = h.userService.Delete(ctx, deleterId.(uint64), uint64(userIdToDelete))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			return api.NewNotFoundError(err, "Not found user to delete")
		case errors.Is(err, service.ErrForbidden):
			return api.NewForbiddenError(err, "Deletion forbidden")
		default:
			return api.NewInternalError(err)
		}
	}
	c.Status(http.StatusNoContent)
	return nil
}
