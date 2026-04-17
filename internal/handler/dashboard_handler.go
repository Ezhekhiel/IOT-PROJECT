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
func (h *DashboardHandler) GetHistory(c *gin.Context) {

	deviceCode := c.Param("device_code")
	rangeType := c.DefaultQuery("range", "1d")

	data, err := h.Service.GetHistory(deviceCode, rangeType)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"device_code": deviceCode,
		"range":       rangeType,
		"data":        data,
	})
}
