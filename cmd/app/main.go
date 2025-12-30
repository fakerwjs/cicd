package main

import (
	"cicd/internal/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", handler.Health)
	r.GET("/hello", handler.Hello)
	r.GET("/test", handler.Test)
	log.Println("server start :8080")
	r.Run(":8080")
}
