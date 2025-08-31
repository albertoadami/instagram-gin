package main

import (
	"github.com/albertoadami/instagram-gin/internal/configuration"
	"github.com/albertoadami/instagram-gin/internal/database"
	"github.com/albertoadami/instagram-gin/internal/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	router := gin.Default()
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	logger.Info("Starting server...")

	config, err := configuration.LoadConfig()
	if err != nil {
		logger.Fatal("Exiting due to config error")
	}

	db, err := database.Connect(&config.Database)
	if err != nil {
		logger.Sugar().Errorw("Failed to connect to database", "error", err)
		logger.Fatal("Exiting due to DB error")
	}
	defer db.Close()

	// handlers
	healthHandler := handlers.NewHealthHandler(db)

	router.GET("/health", healthHandler.HealthCheckHandler)

	router.Run() // listen and serve on 0.0.0.0:8080
}
