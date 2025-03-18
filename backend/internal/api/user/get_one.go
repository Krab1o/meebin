package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Krab1o/meebin/internal/api"
	convUser "github.com/Krab1o/meebin/internal/converter/api/user"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/gin-gonic/gin"
)

//	@Tags			User
//	@Summary		Get user by ID
//	@Schemes		http
//	@Description	Returns user by ID
//	@Accept			json
//	@Produce		json
//	@Security		jwtToken
//	@Param			user_id	path		int	true	"User ID"
//	@Success		200		{object}	dto.BaseUser
//	@Failure		400		{object}	api.Error
//	@Failure		401		{object}	api.Error
//	@Failure		404		{object}	api.Error
//	@Failure		500		{object}	api.Error
//	@Router			/users/{user_id} [get]
func (h *handler) GetUser(c *gin.Context) error {
	ctx := c.Request.Context()
	idUnparsed := c.Param(api.ParamId)
	userId, err := strconv.Atoi(idUnparsed)
	if err != nil {
		return api.NewBadRequestError(err, "Wrong id format")
	}
	if userId < 0 {
		return api.NewBadRequestError(nil, "Wrong id value")
	}
	user, err := h.userService.GetUser(ctx, uint64(userId))
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			return api.NewNotFoundError(err, "User not found")
		default:
			return api.NewInternalError(err)
		}
	}
	user.Id = uint64(userId)
	dtoUser := convUser.UserServiceToDTO(user)
	c.JSON(http.StatusOK, dtoUser)
	return nil
}
