package auth

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api/auth/converter"
	"github.com/Krab1o/meebin/internal/struct/dto"
	"github.com/gin-gonic/gin"
)

// TODO: solidify error handling
// TODO: add credentials validation
func (h *handler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	user := &dto.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	serviceUser := converter.UserDTOToService(user)
	tokens, err := h.authService.Register(ctx, serviceUser)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	dtoTokens := converter.TokensServiceToDTO(tokens)
	c.JSON(http.StatusCreated, token)
}
