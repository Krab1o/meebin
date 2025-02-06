package server

import "github.com/gin-gonic/gin"

func startHttpServer() {
	s := gin.Default()
	s.Run()
}