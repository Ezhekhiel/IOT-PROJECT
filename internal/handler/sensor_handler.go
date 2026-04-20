package handler

import (
	"net/http"

	"github.com/Ezhekhiel/IOT-PROJECT/internal/service"
	"github.com/gin-gonic/gin"
)

type SensorHandler struct {
	Service service.SensorService
}

func (h *SensorHandler) ReceiveSensor(c *gin.Context) {
	var req service.SensorRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.DeviceCode == "" {
		c.JSON(400, gin.H{"error": "device_code is required"})
		return
	}

	if req.Pressure == nil && req.Temperature == nil && req.Timer == nil {
		c.JSON(400, gin.H{"error": "at least one sensor value required"})
		return
	}

	status, err := h.Service.Process(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
