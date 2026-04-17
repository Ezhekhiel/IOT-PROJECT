package service

import (
	"errors"

	"github.com/Ezhekhiel/IOT-PROJECT/internal/model"
	"github.com/Ezhekhiel/IOT-PROJECT/internal/repository"
)

type SensorService struct {
	DeviceRepo      repository.DeviceRepository
	CellRunningRepo repository.CellRunningModelRepository
	StandardRepo    repository.StandardRepository
	SensorRepo      repository.SensorRepository
	AlertRepo       repository.AlertRepository
}

type SensorRequest struct {
	DeviceCode  string   `json:"device_code"`
	Pressure    *float64 `json:"pressure"`
	Temperature *float64 `json:"temperature"`
	Timer       *int     `json:"timer"`
}

func (s *SensorService) Process(data SensorRequest) (string, error) {

	// 1. Find Device
	device, err := s.DeviceRepo.FindByCode(data.DeviceCode)
	if err != nil {
		return "", errors.New("device not found")
	}

	// 2. Find Running Model
	running, err := s.CellRunningRepo.FindLatestByCell(device.CellID)
	if err != nil {
		return "", errors.New("no running model")
	}

	// 3. Find Standard
	standard, err := s.StandardRepo.Find(running.ModelID, device.ProcessID)
	if err != nil {
		return "", errors.New("standard not found")
	}

	status := "OK"
	message := ""

	// 4. VALIDATION LOGIC

	// Pressure
	if data.Pressure != nil && standard.MaxPressure != nil {
		if *data.Pressure > *standard.MaxPressure {
			status = "ALERT"
			message = "Over Pressure"
		}
	}

	if data.Pressure != nil && standard.MinPressure != nil {
		if *data.Pressure < *standard.MinPressure {
			status = "ALERT"
			message = "Low Pressure"
		}
	}

	// Temperature
	if data.Temperature != nil && standard.MaxTemperature != nil {
		if *data.Temperature > *standard.MaxTemperature {
			status = "ALERT"
			message = "Over Temperature"
		}
	}

	if data.Temperature != nil && standard.MinTemperature != nil {
		if *data.Temperature < *standard.MinTemperature {
			status = "ALERT"
			message = "Low Temperature"
		}
	}

	// Timer
	if data.Timer != nil && standard.MaxTimer != nil {
		if *data.Timer > *standard.MaxTimer {
			status = "ALERT"
			message = "Over Timer"
		}
	}

	// 5. SAVE SENSOR
	sensor := model.SensorData{
		DeviceID:    device.ID,
		Pressure:    data.Pressure,
		Temperature: data.Temperature,
		Timer:       data.Timer,
	}

	err = s.SensorRepo.Save(sensor)
	if err != nil {
		return "", err
	}

	// 6. ALERT HANDLING
	if status == "ALERT" {
		alert := model.Alert{
			DeviceID: device.ID,
			Message:  message,
		}
		s.AlertRepo.Create(alert)
	} else {
		s.AlertRepo.ResolveActiveAlert(device.ID)
	}

	return status, nil
}
