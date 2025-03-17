package user

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	convUser "github.com/Krab1o/meebin/internal/converter/api/user"
	"github.com/Krab1o/meebin/internal/model/dto"
	"github.com/Krab1o/meebin/internal/shared"
	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateUser(c *gin.Context) error {
	ctx := c.Request.Context()
	id, ok := c.Get(shared.UserIDJsonName)
	if !ok {
		return api.NewInternalError(nil, "Unable to parse sessionId")
	}
	user := &dto.User{
		Id: id.(uint64),
	}
	err := c.ShouldBindJSON(user)
	if err != nil {
		return api.NewBadRequestError(err, api.ParseValidationErrors(err))
	}
	serviceUser := convUser.UserDTOToService(user)
	newUser, err := h.userService.Update(ctx, serviceUser)
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}
	dtoUser := convUser.UserServiceToDTO(newUser)
	c.JSON(http.StatusOK, dtoUser)
	return nil
}
