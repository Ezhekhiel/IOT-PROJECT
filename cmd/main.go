package main

import (
	"github.com/Ezhekhiel/IOT-PROJECT/internal/config"
	"github.com/Ezhekhiel/IOT-PROJECT/internal/handler"
	"github.com/Ezhekhiel/IOT-PROJECT/internal/repository"
	"github.com/Ezhekhiel/IOT-PROJECT/internal/service"
	"github.com/Ezhekhiel/IOT-PROJECT/middleware"

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
		StandarRepo:     standardRepo,
		CellRunningRepo: cellRepo,
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
	api := r.Group("/api", middleware.APIKeyMiddleware())
	//data input dari IoT
	api.POST("/sensor", sensorHandler.ReceiveSensor)
	//dashboard data
	api.GET("/dashboard/latest/:device_code", dashboardHandler.GetLatest)
	api.GET("/dashboard/latest/", dashboardHandler.GetLatestAll)
	//report data
	api.GET("/dashboard/history/:device_code", dashboardHandler.GetHistory)
	api.GET("/dashboard/history/", dashboardHandler.GetHistoryAll)

	//get data cell running models
	api.GET("/data/", dataHandler.GetModelProcessStandards)
	api.POST("/data/", dataHandler.GetModelProcessStandards)
	api.GET("/data/cell", dataHandler.GetAllCell)
	api.POST("/data/cell", dataHandler.ReceiveCell)
	api.GET("/data/location", dataHandler.GetAllLocation)
	api.POST("/data/location", dataHandler.ReceiveLocation)
	api.GET("/data/model", dataHandler.GetAllModel)
	api.POST("/data/model", dataHandler.ReceiveModel)
	api.GET("/data/process", dataHandler.GetAllProcess)
	api.POST("/data/process", dataHandler.ReceiveProcess)

	api.GET("/alerts/active", alertHandler.GetActive)

	r.Run(":8080")
}
