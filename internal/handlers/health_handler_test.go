package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/albertoadami/instagram-gin/internal/testutil"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandlerSuccess(t *testing.T) {
	context := context.Background()
	container, err := testutil.SetupPostgresWithoutMigrations(context)
	defer container.Teardown(context)
	assert.NoError(t, err)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	healthHandler := NewHealthHandler(container.DB)
	router.GET("/health", healthHandler.HealthCheckHandler)

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestHealthHandlerDBDown(t *testing.T) {
	context := context.Background()
	container, err := testutil.SetupPostgresWithoutMigrations(context)
	defer container.Teardown(context)
	assert.NoError(t, err)

	// Simulate DB down
	container.Teardown(context)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	healthHandler := NewHealthHandler(container.DB)
	router.GET("/health", healthHandler.HealthCheckHandler)

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
