package auth

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	convToken "github.com/Krab1o/meebin/internal/converter/api/token"
	convUser "github.com/Krab1o/meebin/internal/converter/api/user"
	"github.com/Krab1o/meebin/internal/model/dto"
	"github.com/gin-gonic/gin"
)

// TODO: add messages to API errors
func (h *handler) Login(c *gin.Context) error {
	ctx := c.Request.Context()
	userCreds := &dto.Creds{}
	err := c.ShouldBindJSON(userCreds)
	if err != nil {
		return api.NewBadRequestError(err, api.ParseValidationErrors(err))
	}
	serviceCreds := convUser.CredsDTOToService(userCreds)
	tokens, err := h.authService.Login(ctx, serviceCreds)
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}
	dtoTokens := convToken.TokensServiceToDTO(tokens)
	c.JSON(http.StatusOK, dtoTokens)
	return nil
}
