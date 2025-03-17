package auth

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	convUser "github.com/Krab1o/meebin/internal/converter/api/new_user"
	convToken "github.com/Krab1o/meebin/internal/converter/api/token"
	"github.com/Krab1o/meebin/internal/model/dto"
	"github.com/gin-gonic/gin"
)

const (
	invalidCredentialsMessage = "Invalid credentials"
)

func (h *handler) Register(c *gin.Context) error {
	ctx := c.Request.Context()
	newUser := &dto.NewUser{}
	err := c.ShouldBindJSON(newUser)
	if err != nil {
		return api.NewBadRequestError(err, api.ParseValidationErrors(err))
	}

	serviceUser := convUser.NewUserDTOToService(newUser)
	tokens, err := h.authService.Register(ctx, serviceUser)
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}
	dtoTokens := convToken.TokensServiceToDTO(tokens)
	c.JSON(http.StatusCreated, dtoTokens)
	return nil
}
