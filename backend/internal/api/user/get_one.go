package user

import (
	"net/http"
	"strconv"

	"github.com/Krab1o/meebin/internal/api"
	convUser "github.com/Krab1o/meebin/internal/converter/api/user"
	"github.com/gin-gonic/gin"
)

func (h *handler) GetUser(c *gin.Context) error {
	ctx := c.Request.Context()
	idUnparsed := c.Param(api.ParamId)
	userId, err := strconv.Atoi(idUnparsed)
	if err != nil {
		return api.NewBadRequestError(err, "Wrong id format")
	}
	if userId < 0 {
		return api.NewBadRequestError(err, "Wrong id value")
	}
	user, err := h.userService.GetUser(ctx, uint64(userId))
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}
	user.Id = uint64(userId)
	dtoUser := convUser.UserServiceToDTO(user)
	c.JSON(http.StatusOK, dtoUser)
	return nil
}
