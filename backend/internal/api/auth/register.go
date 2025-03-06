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

// TODO: add credentials validation
// TODO: add validation for personal data and creds fields
func (h *handler) Register(c *gin.Context) error {
	ctx := c.Request.Context()
	newUser := &dto.RequestCreateUser{}
	err := c.ShouldBindJSON(newUser)
	if err != nil {
		return api.NewBadRequestError(err, "Cannot process entity")
	}
	if errs := api.ValidateStruct(*h.validate, newUser); len(errs) > 0 {
		return api.NewBadRequestError(nil, errs)
	}
	serviceUser := converter.NewUserDTOToService(newUser)
	tokens, err := h.authService.Register(ctx, serviceUser)
	if err != nil {
		return api.ErrorServiceToAPI(err, nil)
	}
	dtoTokens := converter.ResponseTokensServiceToDTO(tokens)
	c.JSON(http.StatusCreated, dtoTokens)
	return nil
}
