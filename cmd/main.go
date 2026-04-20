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

	dashboardService := service.DashboardService{
		SensorRepo: sensorRepo,
		DeviceRepo: deviceRepo,
	}
	alertService := service.AlertService{
		AlertRepo: alertRepo, // Sesuaikan dengan field yang ada di AlertService
	}
	dataService := service.DataService{
		DataRepo: standardRepo,
	}

	// handler
	sensorHandler := handler.SensorHandler{
		Service: sensorService,
	}

	dashboardHandler := handler.DashboardHandler{
		Service: dashboardService,
	}
	alertHandler := handler.AlertHandler{
		Service: alertService,
	}
	dataHandler := handler.DataHandler{
		Service: dataService,
	}

	r := gin.Default()
	//data input dari IoT
	r.POST("/api/sensor", sensorHandler.ReceiveSensor)
	//dashboard data
	r.GET("/api/dashboard/latest/:device_code", dashboardHandler.GetLatest)
	r.GET("/api/dashboard/latest/", dashboardHandler.GetLatestAll)
	//report data
	r.GET("/api/dashboard/history/:device_code", dashboardHandler.GetHistory)
	r.GET("/api/dashboard/history/", dashboardHandler.GetHistoryAll)

	//get data cell running models
	r.GET("/api/data/cell/running_model", dataHandler.GetModelProcessStandards)

	r.GET("/api/alerts/active", alertHandler.GetActive)

	r.Run(":8080")
}
