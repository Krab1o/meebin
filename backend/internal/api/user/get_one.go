package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *handler) GetUser(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}