package event

import (
	"errors"
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	converter "github.com/Krab1o/meebin/internal/converter/api/event/create"
	"github.com/Krab1o/meebin/internal/model/event/dto"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/gin-gonic/gin"
)

// @Tags			Event
// @Summary		Creates event
// @Schemes		http
// @Description	Creates a new event in the system
// @Accept			json
// @Produce		json
// @Security		jwtToken
// @Param			EventData	body		dto.NewEvent	true	"New Event Info"
// @Success		201			{object}	dto.IdResponse
// @Failure		400			{object}	api.Error
// @Failure		409			{object}	api.Error
// @Failure		500			{object}	api.Error
// @Router			/events [post]
func (h *Handler) Create(c *gin.Context) error {
	ctx := c.Request.Context()
	callerId, ok := c.Get(shared.UserIDJsonName)
	if !ok {
		return api.NewInternalError(nil)
	}
	event := &dto.NewEvent{}
	err := c.ShouldBindJSON(event)
	if err != nil {
		return api.NewBadRequestError(err, api.ParseValidationErrors(err))
	}

	serviceEvent := converter.NewEventServiceToDTO(event)
	eventId, err := h.eventService.Create(ctx, callerId.(uint64), serviceEvent)
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
