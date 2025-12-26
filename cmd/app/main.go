package main

import (
	"log"

	"cicd/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", handler.Health)
	r.GET("/hello", handler.Hello)

	log.Println("server start :8080")
	r.Run(":8080")
}
