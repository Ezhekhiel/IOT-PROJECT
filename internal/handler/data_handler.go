package handler

import (
	"net/http"

	"github.com/Ezhekhiel/IOT-PROJECT/internal/service"
	"github.com/gin-gonic/gin"
)

type DataHandler struct {
	Service service.DataService
}

func (h *DataHandler) GetModelProcessStandards(c *gin.Context) {
	data, err := h.Service.GetModelProcessStandards()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
