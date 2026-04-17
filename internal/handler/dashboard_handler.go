package handler

import (
	"net/http"

	"github.com/Ezhekhiel/IOT-PROJECT/internal/service"
	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	Service service.DashboardService
}

func (h *DashboardHandler) GetLatest(c *gin.Context) {

	deviceCode := c.Param("device_code")

	data, err := h.Service.GetLatest(deviceCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
