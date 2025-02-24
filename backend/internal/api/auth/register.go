package auth

import (
	"github.com/Krab1o/meebin/internal/api/auth/converter"
	"github.com/Krab1o/meebin/internal/api/auth/dto"
	"github.com/gin-gonic/gin"
)

func (h *handler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	creds := &dto.UserCreds{}
	serviceCreds := converter.CredsDTOToService(creds)
	h.authService.Register(ctx, *serviceCreds)
}
