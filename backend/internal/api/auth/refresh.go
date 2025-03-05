package auth

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/struct/dto"
	"github.com/gin-gonic/gin"
)

func (h *handler) Refresh(c *gin.Context) error {
	ctx := c.Request.Context()
	token := &dto.Token{}
	err := c.ShouldBindJSON(token)
	if err != nil {
		return api.NewBadRequestError(err, nil)
	}
	newRefreshToken, err := h.authService.Refresh(ctx, token.Refresh)
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}
	c.JSON(http.StatusOK, &dto.Token{Refresh: newRefreshToken})
	return nil
}
