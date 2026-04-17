package handler

import (
	"github.com/Ezhekhiel/IOT-PROJECT/internal/service"
	"github.com/gin-gonic/gin"
)

type AlertHandler struct {
	Service service.AlertService
}

func (h *AlertHandler) GetActive(c *gin.Context) {

	data, err := h.Service.GetActiveAlerts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": data,
	})
}
