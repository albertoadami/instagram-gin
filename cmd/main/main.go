package main

import (
	"github.com/albertoadami/instagram-gin/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", handlers.HealthHandler)

	router.Run() // listen and serve on 0.0.0.0:8080

}
