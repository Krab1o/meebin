package event

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Krab1o/meebin/internal/api"
	converter "github.com/Krab1o/meebin/internal/converter/api/event/base"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/gin-gonic/gin"
)

// @Tags			Event
// @Summary		Get event by ID
// @Schemes		http
// @Description	Returns event by ID
// @Accept			json
// @Produce		json
// @Security		jwtToken
// @Param			event_id	path		int	true	"Event ID"
// @Success		200		{object}	dto.BaseEvent
// @Failure		400		{object}	api.Error
// @Failure		401		{object}	api.Error
// @Failure		404		{object}	api.Error
// @Failure		500		{object}	api.Error
// @Router			/events/{event_id} [get]
func (h *Handler) Get(c *gin.Context) error {
	ctx := c.Request.Context()
	idUnparsed := c.Param(api.ParamId)
	eventId, err := strconv.Atoi(idUnparsed)
	if err != nil {
		return api.NewBadRequestError(err, "Wrong id format")
	}
	if eventId < 0 {
		return api.NewBadRequestError(nil, "Wrong id value")
	}
	event, err := h.eventService.Get(ctx, uint64(eventId))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			return api.NewNotFoundError(err, "User not found")
		default:
			return api.NewInternalError(err)
		}
	}
	event.Id = uint64(eventId)
	dtoEvent := converter.BaseEventServiceToDTO(event)
	c.JSON(http.StatusOK, dtoEvent)
	return nil
}
