package service

import (
	"time"

	"github.com/Ezhekhiel/IOT-PROJECT/internal/repository"
)

type DashboardService struct {
	SensorRepo repository.SensorRepository
	DeviceRepo repository.DeviceRepository
}
type HistoryAggregated struct {
	TimeGroup   time.Time `gorm:"column:time_group"`
	Pressure    *float64  `gorm:"column:pressure"`
	Temperature *float64  `gorm:"column:temperature"`
	Timer       *float64  `gorm:"column:timer"`
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
func (s *DashboardService) GetLatestFromAll() (interface{}, error) {

	data, err := s.SensorRepo.GetLatestAllDevices()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}

	for _, d := range data {
		result = append(result, map[string]interface{}{
			"device_id":   d.DeviceID,
			"pressure":    d.Pressure,
			"temperature": d.Temperature,
			"timer":       d.Timer,
			"time":        d.CreatedAt,
		})
	}

	return result, nil
}
func (s *DashboardService) GetHistory(deviceCode string, rangeType string) (interface{}, error) {

	device, err := s.DeviceRepo.FindByCode(deviceCode)
	if err != nil {
		return nil, err
	}

	data, err := s.SensorRepo.GetHistory(device.ID, rangeType)
	if err != nil {
		return nil, err
	}

	// format untuk chart
	var result []map[string]interface{}

	for _, d := range data {
		result = append(result, map[string]interface{}{
			"time":        d.CreatedAt,
			"pressure":    d.Pressure,
			"temperature": d.Temperature,
			"timer":       d.Timer,
		})
	}

	return result, nil
}
func (s *DashboardService) GetHistoryFromAll(rangeType string) (interface{}, error) {

	data, err := s.SensorRepo.GetHistoryAggregated(rangeType)
	if err != nil {
		return nil, err
	}

	// format untuk chart
	var result []map[string]interface{}
	for _, d := range data {
		result = append(result, map[string]interface{}{
			"time":        d.TimeGroup,
			"pressure":    d.Pressure,
			"temperature": d.Temperature,
			"timer":       d.Timer,
			"device_code": d.DeviceCode,
		})
	}
	return result, nil
}
