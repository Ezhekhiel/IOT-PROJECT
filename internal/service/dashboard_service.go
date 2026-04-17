package service

import "github.com/Ezhekhiel/IOT-PROJECT/internal/repository"

type DashboardService struct {
	SensorRepo repository.SensorRepository
	DeviceRepo repository.DeviceRepository
}

func (s *DashboardService) GetLatest(deviceCode string) (interface{}, error) {

	device, err := s.DeviceRepo.FindByCode(deviceCode)
	if err != nil {
		return nil, err
	}

	data, err := s.SensorRepo.GetLatestByDevice(device.ID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"device_code": deviceCode,
		"pressure":    data.Pressure,
		"temperature": data.Temperature,
		"timer":       data.Timer,
		"time":        data.CreatedAt,
	}, nil
}
