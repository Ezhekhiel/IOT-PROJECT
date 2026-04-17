package main

import (
	"github.com/Ezhekhiel/IOT-PROJECT/internal/config"
	"github.com/Ezhekhiel/IOT-PROJECT/internal/handler"
	"github.com/Ezhekhiel/IOT-PROJECT/internal/repository"
	"github.com/Ezhekhiel/IOT-PROJECT/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	// repo
	deviceRepo := repository.DeviceRepository{DB: db}
	cellRepo := repository.CellRunningModelRepository{DB: db}
	standardRepo := repository.StandardRepository{DB: db}
	sensorRepo := repository.SensorRepository{DB: db}
	alertRepo := repository.AlertRepository{DB: db}

	// service
	sensorService := service.SensorService{
		DeviceRepo:      deviceRepo,
		CellRunningRepo: cellRepo,
		StandardRepo:    standardRepo,
		SensorRepo:      sensorRepo,
		AlertRepo:       alertRepo,
	}

	// handler
	sensorHandler := handler.SensorHandler{
		Service: sensorService,
	}
	dashboardService := service.DashboardService{
		SensorRepo: sensorRepo,
		DeviceRepo: deviceRepo,
	}

	dashboardHandler := handler.DashboardHandler{
		Service: dashboardService,
	}

	r := gin.Default()

	r.POST("/api/sensor", sensorHandler.ReceiveSensor)
	r.GET("/api/dashboard/latest/:device_code", dashboardHandler.GetLatest)

	r.Run(":8080")
}
