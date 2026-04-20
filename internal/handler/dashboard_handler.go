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
	if deviceCode == "" {
		c.JSON(400, gin.H{"error": "device_code is required"})
		return
	}

	data, err := h.Service.GetLatest(deviceCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
func (h *DashboardHandler) GetLatestAll(c *gin.Context) {

	data, err := h.Service.GetLatestFromAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
func (h *DashboardHandler) GetHistory(c *gin.Context) {

	deviceCode := c.Param("device_code")
	rangeType := c.DefaultQuery("range", "1d")

	if deviceCode == "" {
		c.JSON(400, gin.H{"error": "device_code is required"})
		return
	}

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
func (h *DashboardHandler) GetHistoryAll(c *gin.Context) {

	rangeType := c.DefaultQuery("range", "1d")

	data, err := h.Service.GetHistoryFromAll(rangeType)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"range": rangeType,
		"data":  data,
	})
}
