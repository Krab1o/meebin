package event

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/gin-gonic/gin"
)

// @Tags			Event
// @Summary		Deletes event
// @Schemes		http
// @Description	Delete an event from the system
// @Accept			json
// @Produce		json
// @Security		jwtToken
// @Param			event_id	path	int	true	"Event ID"
// @Success		204
// @Failure		400			{object}	api.Error
// @Failure		401			{object}	api.Error
// @Failure		409			{object}	api.Error
// @Failure		500			{object}	api.Error
// @Router			/events/{event_id} [delete]
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
	eventId, err := strconv.Atoi(paramId)
	if err != nil {
		return api.NewInternalError(nil, "Unable to convert query param")
	}
	err = h.eventService.Delete(ctx, deleterId.(uint64), uint64(eventId))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			return api.NewNotFoundError(err, "Not found event to delete")
		default:
			return api.NewInternalError(err)
		}
	}
	c.Status(http.StatusNoContent)
	return nil
}
