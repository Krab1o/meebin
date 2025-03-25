package auth

import (
	"errors"
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	convToken "github.com/Krab1o/meebin/internal/converter/api/token"
	convUser "github.com/Krab1o/meebin/internal/converter/api/user/login"
	"github.com/Krab1o/meebin/internal/model/user/dto"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/gin-gonic/gin"
)

// TODO: add login DTO
// TODO: add separated login credentials

// @Tags			Auth
// @Summary		Login
// @Schemes		http
// @Description	Creates new pair of refresh-access tokens based on your credentials
// @Accept			json
// @Produce		json
// @Param			LoginCreds	body		dto.LoginCreds	true	"Login info"
// @Success		200			{object}	dto.ResponseTokens
// @Failure		400			{object}	api.Error
// @Failure		401			{object}	api.Error
// @Failure		404			{object}	api.Error
// @Failure		500			{object}	api.Error
// @Router			/auth/login [post]
func (h *Handler) Login(c *gin.Context) error {
	ctx := c.Request.Context()
	userCreds := &dto.LoginCreds{}
	err := c.ShouldBindJSON(userCreds)
	if err != nil {
		return api.NewBadRequestError(err, api.ParseValidationErrors(err))
	}
	serviceCreds := convUser.LoginCredsDTOToService(userCreds)
	tokens, err := h.authService.Login(ctx, serviceCreds)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			return api.NewNotFoundError(err, "User not found")
		case errors.Is(err, service.ErrUnautorized):
			return api.NewUnauthorizedError(err, "Wrong password")
		default:
			return api.NewInternalError(err)
		}
	}
	dtoTokens := convToken.TokensServiceToDTO(tokens)
	c.JSON(http.StatusOK, dtoTokens)
	return nil
}
