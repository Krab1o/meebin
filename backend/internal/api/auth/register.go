package auth

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/api/auth/converter"
	"github.com/Krab1o/meebin/internal/struct/dto"
	"github.com/gin-gonic/gin"
)

const (
	invalidCredentialsMessage = "Invalid credentials"
)

// TODO: add messages to API errors
// TODO: add credentials validation
// TODO: add validation for personal data and creds fields
func (h *handler) Register(c *gin.Context) error {
	ctx := c.Request.Context()
	newUser := &dto.NewUser{}
	err := c.ShouldBindJSON(newUser)
	if err != nil {
		return api.ErrorServiceToAPI("", err)
	}
	serviceUser := converter.NewUserDTOToService(newUser)
	tokens, err := h.authService.Register(ctx, serviceUser)
	if err != nil {
		return api.ErrorServiceToAPI("", err)
	}
	dtoTokens := converter.TokensServiceToDTO(tokens)
	c.JSON(http.StatusCreated, dtoTokens)
	return nil
}
