package event

import (
	"errors"
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	converter "github.com/Krab1o/meebin/internal/converter/api/event/create"
	"github.com/Krab1o/meebin/internal/model/event/dto"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/gin-gonic/gin"
)

// TODO: create DTOs
// TODO: create handlers
func (h *Handler) Create(c *gin.Context) error {
	ctx := c.Request.Context()
	event := &dto.NewEvent{}
	err := c.ShouldBindJSON(event)
	if err != nil {
		return api.NewBadRequestError(err, api.ParseValidationErrors(err))
	}

	serviceEvent := converter.NewEventServiceToDTO(event)
	eventId, err := h.eventService.Create(ctx, serviceEvent)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrDuplicate):
			return api.NewDuplicateError(err, "Event already exists")
		default:
			return api.NewInternalError(err, "Internal Error")
		}
	}
	c.JSON(http.StatusCreated, dto.IdResponse{Id: eventId})
	return nil
}
