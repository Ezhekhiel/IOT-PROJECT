package service

import (
	"github.com/Ezhekhiel/IOT-PROJECT/internal/model"
	"github.com/Ezhekhiel/IOT-PROJECT/internal/repository"
)

type AlertService struct {
	AlertRepo repository.AlertRepository
}

func (s *AlertService) GetActiveAlerts() ([]model.Alert, error) {
	return s.AlertRepo.GetActive()
}
func (s *AlertService) Create(deviceID int, message string) error {

	alert := model.Alert{
		DeviceID: deviceID,
		Message:  message,
		Status:   "ACTIVE",
	}

	return s.AlertRepo.Create(alert)
}
