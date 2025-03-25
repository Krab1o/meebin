package event

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	converter "github.com/Krab1o/meebin/internal/converter/api/event/base"
	"github.com/Krab1o/meebin/internal/model/event/dto"
	"github.com/gin-gonic/gin"
)

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
