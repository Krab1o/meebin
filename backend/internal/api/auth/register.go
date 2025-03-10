package auth

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/api/auth/converter"
	"github.com/Krab1o/meebin/internal/model/dto"
	"github.com/gin-gonic/gin"
)

const (
	invalidCredentialsMessage = "Invalid credentials"
)

func (h *handler) Register(c *gin.Context) error {
	ctx := c.Request.Context()
	newUser := &dto.RequestCreateUser{}
	err := c.ShouldBindJSON(newUser)
	if err != nil {
		return api.NewBadRequestError(err, api.ParseValidationErrors(err))
	}
	serviceUser := converter.RequestUserDTOToService(newUser)
	tokens, err := h.authService.Register(ctx, serviceUser)
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}
	dtoTokens := converter.ResponseTokensServiceToDTO(tokens)
	c.JSON(http.StatusCreated, dtoTokens)
	return nil
}
