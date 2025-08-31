package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupPostgres(t *testing.T) (host string, port string, terminate func()) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15",
		ExposedPorts: []string{"5432/tcp"}, // Docker will map this to a random host port
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)

	host, err = postgres.Host(ctx)
	assert.NoError(t, err)
	portObj, err := postgres.MappedPort(ctx, "5432")
	assert.NoError(t, err)

	terminate = func() {
		_ = postgres.Terminate(ctx)
	}
	return host, portObj.Port(), terminate
}

func TestHealthHandlerSuccess(t *testing.T) {
	host, port, terminate := setupPostgres(t)
	defer terminate()

	dsn := fmt.Sprintf("host=%s port=%s user=testuser password=testpass dbname=testdb sslmode=disable", host, port)
	db, err := sqlx.Connect("postgres", dsn)
	assert.NoError(t, err)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	healthHandler := NewHealthHandler(db)
	router.GET("/health", healthHandler.HealthCheckHandler)

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestHealthHandlerDBDown(t *testing.T) {
	// Start PostgreSQL container
	host, port, terminate := setupPostgres(t)
	defer terminate()

	dsn := fmt.Sprintf("host=%s port=%s user=testuser password=testpass dbname=testdb sslmode=disable", host, port)
	db, err := sqlx.Connect("postgres", dsn)
	assert.NoError(t, err)

	// Simulate DB down
	db.Close()

	gin.SetMode(gin.TestMode)
	router := gin.New()
	healthHandler := NewHealthHandler(db)
	router.GET("/health", healthHandler.HealthCheckHandler)

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
