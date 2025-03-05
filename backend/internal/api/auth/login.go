package auth

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/api/auth/converter"
	"github.com/Krab1o/meebin/internal/struct/dto"
	"github.com/gin-gonic/gin"
)

// TODO: add messages to API errors
func (h *handler) Login(c *gin.Context) error {
	ctx := c.Request.Context()
	userCreds := &dto.Creds{}
	err := c.ShouldBindJSON(userCreds)
	if err != nil {
		return api.NewBadRequestError(err, "Bad request specified")
	}
	if userCreds.Email == "" || userCreds.Username == "" {
		return api.NewBadRequestError(err, "Empty password or username")
	}
	serviceCreds := converter.CredsDTOToService(userCreds)
	tokens, err := h.authService.Login(ctx, serviceCreds)
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}
	dtoTokens := converter.TokensServiceToDTO(tokens)
	c.JSON(http.StatusOK, dtoTokens)
	return nil
}
