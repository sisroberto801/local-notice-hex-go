package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
}

func (h *HealthHandler) Health(c *gin.Context) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "local-notice-hex-go",
		Version:   "1.0.0",
	}

	c.JSON(http.StatusOK, response)
}

func (h *HealthHandler) Ready(c *gin.Context) {
	response := HealthResponse{
		Status:    "ready",
		Timestamp: time.Now(),
		Service:   "local-notice-hex-go",
		Version:   "1.0.0",
	}

	c.JSON(http.StatusOK, response)
}
