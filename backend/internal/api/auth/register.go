package auth

import (
	"errors"
	"net/http"

	"github.com/Krab1o/meebin/internal/api"
	convToken "github.com/Krab1o/meebin/internal/converter/api/token"
	convUser "github.com/Krab1o/meebin/internal/converter/api/user/register"
	"github.com/Krab1o/meebin/internal/model/user/dto"
	"github.com/Krab1o/meebin/internal/service"
	"github.com/gin-gonic/gin"
)

const (
	invalidCredentialsMessage = "Invalid credentials"
)

// @Tags			Auth
// @Summary		Registration
// @Schemes		http
// @Description	Registers a new user in service
// @Accept			json
// @Produce		json
// @Param			UserData	body		dto.RegisterUser	true	"New User Info"
// @Success		201			{object}	dto.ResponseTokens
// @Failure		400			{object}	api.Error
// @Failure		409			{object}	api.Error
// @Failure		500			{object}	api.Error
// @Router			/auth/register [post]
func (h *Handler) Register(c *gin.Context) error {
	ctx := c.Request.Context()
	newUser := &dto.RegisterUser{}
	err := c.ShouldBindJSON(newUser)
	if err != nil {
		return api.NewBadRequestError(err, api.ParseValidationErrors(err))
	}

	serviceUser := convUser.NewUserDTOToService(newUser)
	tokens, err := h.authService.Register(ctx, serviceUser)
	if err != nil {
		// log.Println("+++")
		// fmt.Println(errors.Is(err, service.ErrDuplicate))
		// log.Println("+++")
		switch {
		case errors.Is(err, service.ErrDuplicate):
			return api.NewDuplicateError(err, "User already exists")
		default:
			return api.NewInternalError(err, "Internal Error")
		}
	}
	dtoTokens := convToken.TokensServiceToDTO(tokens)
	c.JSON(http.StatusCreated, dtoTokens)
	return nil
}
