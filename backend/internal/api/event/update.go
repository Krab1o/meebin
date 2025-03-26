package event

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Krab1o/meebin/internal/api"
	convBase "github.com/Krab1o/meebin/internal/converter/api/event/base"
	convUpdate "github.com/Krab1o/meebin/internal/converter/api/event/update"
	"github.com/Krab1o/meebin/internal/model/event/dto"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/gin-gonic/gin"
)

// @Tags			Event
// @Summary		Updates event
// @Schemes		http
// @Description	Updates event's fields specified in the body.
// @Description	Redundant fields ignored.
// @Description If the field not specified, it is not updated
// @Accept			json
// @Produce		json
// @Security		jwtToken
// @Param			event_id	path		int	true	"Updated event ID"
// @Param UpdatedEntity body dto.UpdateEvent true "Updated event fields"
// @Success		200		{object}	dto.BaseEvent
// @Failure		400		{object}	api.Error
// @Failure		401		{object}	api.Error
// @Failure		403		{object}	api.Error
// @Failure		404		{object}	api.Error
// @Failure		500		{object}	api.Error
// @Router			/events/{event_id} [patch]
func (h *Handler) Update(c *gin.Context) error {
	ctx := c.Request.Context()
	updaterId, ok := c.Get(shared.UserIDJsonName)
	if !ok {
		return api.NewInternalError(nil, "Unable to parse userId")
	}
	paramId, ok := c.Params.Get(api.ParamId)
	if !ok {
		return api.NewInternalError(nil, "Unable to parse query param")
	}
	eventId, err := strconv.Atoi(paramId)
	if err != nil {
		return api.NewInternalError(nil, "Unable to convert query param")
	}
	event := &dto.UpdateEvent{
		Id: uint64(eventId),
	}

	err = c.ShouldBindJSON(event)
	if err != nil {
		return api.NewBadRequestError(err, api.ParseValidationErrors(err))
	}
	serviceEvent := convUpdate.UpdatedEventDTOToService(event)
	updatedEvent, err := h.eventService.Update(ctx, updaterId.(uint64), serviceEvent)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoUpdate):
			return api.NewBadRequestError(err, "No fields to update")
		case errors.Is(err, service.ErrNotFound):
			return api.NewNotFoundError(err, "Event not found")
		case errors.Is(err, service.ErrForbidden):
			return api.NewForbiddenError(err, "Forbidden to update event")
		default:
			return api.NewInternalError(err)
		}
	}
	dtoEvent := convBase.BaseEventServiceToDTO(updatedEvent)
	c.JSON(http.StatusOK, dtoEvent)
	return nil
}
