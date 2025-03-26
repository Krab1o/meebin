package event

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	converter "github.com/Krab1o/meebin/internal/converter/api/event/base"
	"github.com/Krab1o/meebin/internal/model/event/dto"
	"github.com/gin-gonic/gin"
)

//TODO: change returning status for events in the system

// @Tags			Event
// @Summary		Listing events
// @Schemes		http
// @Description	Returns an array of events
// @Accept			json
// @Produce		json
// @Security		jwtToken
// @Success		200	{array}		dto.BaseEvent
// @Failure		401	{object}	api.Error
// @Failure		500	{object}	api.Error
// @Router			/events [get]
func (h *Handler) List(c *gin.Context) error {
	ctx := c.Request.Context()
	events, err := h.eventService.List(ctx)
	if err != nil {
		switch {
		default:
			return api.NewInternalError(err)
		}
	}
	dtoEvents := make([]dto.BaseEvent, len(events))
	for i, event := range events {
		dtoEvents[i] = *converter.BaseEventServiceToDTO(&event)
	}

	c.JSON(http.StatusOK, dtoEvents)
	return nil
}
