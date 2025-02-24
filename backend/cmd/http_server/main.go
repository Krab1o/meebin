package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	s := gin.Default()
	SetupRoutes(s)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	StartHttpServer()
}
