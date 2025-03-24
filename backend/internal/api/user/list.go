package user

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	convUser "github.com/Krab1o/meebin/internal/converter/api/user"
	"github.com/Krab1o/meebin/internal/model/user/dto"
	"github.com/gin-gonic/gin"
)

// @Tags			User
// @Summary		Listing users
// @Schemes		http
// @Description	Returns an array of users
// @Accept			json
// @Produce		json
// @Security		jwtToken
// @Success		200	{array}		dto.BaseUser
// @Failure		401	{object}	api.Error
// @Failure		500	{object}	api.Error
// @Router			/users [get]
func (h *Handler) ListUser(c *gin.Context) error {
	ctx := c.Request.Context()
	users, err := h.userService.ListUser(ctx)
	if err != nil {
		switch {
		default:
			return api.NewInternalError(err)
		}
	}
	dtoUsers := make([]dto.BaseUser, len(users))
	for i, user := range users {
		dtoUsers[i] = *convUser.UserServiceToDTO(&user)
	}

	c.JSON(http.StatusOK, dtoUsers)
	return nil
}
