package user

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	convUser "github.com/Krab1o/meebin/internal/converter/api/user"
	"github.com/Krab1o/meebin/internal/model/dto"
	"github.com/gin-gonic/gin"
)

func (h *handler) ListUser(c *gin.Context) error {
	ctx := c.Request.Context()
	users, err := h.userService.ListUser(ctx)
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}
	dtoUsers := make([]dto.User, len(users))
	for i, user := range users {
		dtoUsers[i] = *convUser.UserServiceToDTO(&user)
	}

	c.JSON(http.StatusOK, dtoUsers)
	return nil
}
