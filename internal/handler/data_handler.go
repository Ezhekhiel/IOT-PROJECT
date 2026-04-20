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
func (h *DataHandler) GetAllCell(c *gin.Context) {
	data, err := h.Service.GetAllCell()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
func (h *DataHandler) ReceiveCell(c *gin.Context) {
	var req service.CellRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Name == "" {
		c.JSON(400, gin.H{"error": "Name is required"})
		return
	}
	if req.LocationId == nil {
		c.JSON(400, gin.H{"error": "Location is required"})
		return
	}
	if req.TargetPerHour == nil {
		c.JSON(400, gin.H{"error": "Target is required"})
		return
	}

	status, err := h.Service.AddCell(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
func (h *DataHandler) GetAllLocation(c *gin.Context) {
	data, err := h.Service.GetAllLocation(0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
func (h *DataHandler) ReceiveLocation(c *gin.Context) {
	var req service.LocationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Name == "" {
		c.JSON(400, gin.H{"error": "Name is required"})
		return
	}
	status, err := h.Service.AddLocation(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
func (h *DataHandler) GetAllModel(c *gin.Context) {
	data, err := h.Service.GetAllModel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
func (h *DataHandler) ReceiveModel(c *gin.Context) {
	var req service.ModelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Name == "" {
		c.JSON(400, gin.H{"error": "Name is required"})
		return
	}
	status, err := h.Service.AddModel(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
func (h *DataHandler) GetAllProcess(c *gin.Context) {
	data, err := h.Service.GetAllProcess()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
func (h *DataHandler) ReceiveProcess(c *gin.Context) {
	var req service.ProcessRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Name == "" {
		c.JSON(400, gin.H{"error": "Name is required"})
		return
	}
	status, err := h.Service.AddProcess(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
