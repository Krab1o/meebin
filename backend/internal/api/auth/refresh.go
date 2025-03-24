package auth

import (
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	"github.com/Krab1o/meebin/internal/model/user/dto"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/gin-gonic/gin"
)

// @Tags			Auth
// @Summary		Refresh
// @Schemes		http
// @Description	Returns new access token based on your refresh token
// @Accept			json
// @Produce		json
// @Param			RefreshToken	body		dto.RefreshToken	true	"User's Refresh Token"
// @Success		200				{object}	dto.AccessToken
// @Failure		400				{object}	api.Error
// @Failure		401				{object}	api.Error
// @Failure		404				{object}	api.Error
// @Failure		500				{object}	api.Error
// @Router			/auth/refresh [post]
func (h *Handler) Refresh(c *gin.Context) error {
	ctx := c.Request.Context()
	token := &dto.RefreshToken{}
	err := c.ShouldBindJSON(token)
	if err != nil {
		return api.NewBadRequestError(err, nil)
	}
	newAccessToken, err := h.authService.Refresh(ctx, token.Refresh)
	if err != nil {
		switch err {
		case service.ErrUnautorized:
			api.NewUnauthorizedError(err, "Invalid refresh token")
		case service.ErrNotFound:
			api.NewNotFoundError(err, "Session not found")
		default:
			return api.NewInternalError(err)
		}
	}
	c.JSON(http.StatusOK, &dto.AccessToken{Access: newAccessToken})
	return nil
}
