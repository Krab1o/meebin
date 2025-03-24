package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Krab1o/meebin/internal/api"
	convUser "github.com/Krab1o/meebin/internal/converter/api/updated_user"
	"github.com/Krab1o/meebin/internal/model/user/dto"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/gin-gonic/gin"
)

// @Tags			User
// @Summary		Update user
// @Schemes		http
// @Description	Updates user's fields specified in the body.
// @Description	Redundant fields ignored.
// @Description If the field not specified, it is not updated
// @Accept			json
// @Produce		json
// @Security		jwtToken
// @Param			user_id	path		int	true	"Updated user ID"
// @Param UpdatedEntity body dto.UpdatedUser true "Updated user fields"
// @Success		200		{object}	dto.UpdatedUser
// @Failure		400		{object}	api.Error
// @Failure		401		{object}	api.Error
// @Failure		403		{object}	api.Error
// @Failure		404		{object}	api.Error
// @Failure		500		{object}	api.Error
// @Router			/users/{user_id} [patch]
func (h *Handler) UpdateUser(c *gin.Context) error {
	ctx := c.Request.Context()
	userIdUpdating, ok := c.Get(shared.UserIDJsonName)
	if !ok {
		return api.NewInternalError(nil, "Unable to parse userId")
	}
	paramId, ok := c.Params.Get(api.ParamId)
	if !ok {
		return api.NewInternalError(nil, "Unable to parse query param")
	}
	userIdToUpdate, err := strconv.Atoi(paramId)
	if err != nil {
		return api.NewInternalError(nil, "Unable to convert query param")
	}

	user := &dto.UpdatedUser{
		Id: userIdUpdating.(uint64),
	}
	err = c.ShouldBindJSON(user)
	if err != nil {
		return api.NewBadRequestError(err, api.ParseValidationErrors(err))
	}
	serviceUser := convUser.UpdatedUserDTOToService(user)

	newUser, err := h.userService.Update(ctx, uint64(userIdToUpdate), serviceUser)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNoUpdate):
			return api.NewBadRequestError(err, "No fields to update")
		case errors.Is(err, service.ErrNotFound):
			return api.NewNotFoundError(err, "User not found")
		case errors.Is(err, service.ErrForbidden):
			return api.NewForbiddenError(err, "Forbidden to update user")
		default:
			return api.NewInternalError(err)
		}
	}

	dtoUser := convUser.UpdatedUserServiceToDTO(newUser)
	c.JSON(http.StatusOK, dtoUser)
	return nil
}
