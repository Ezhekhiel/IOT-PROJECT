package repository

import (
	"github.com/Ezhekhiel/IOT-PROJECT/internal/model"

	"gorm.io/gorm"
)

type SensorRepository struct {
	DB *gorm.DB
}

func (r *SensorRepository) Save(data model.SensorData) error {
	return r.DB.Create(&data).Error
}
func (r *SensorRepository) GetLatestByDevice(deviceID int) (model.SensorData, error) {
	var data model.SensorData

	err := r.DB.
		Where("device_id = ?", deviceID).
		Order("created_at DESC").
		First(&data).Error

	return data, err
}
func (r *SensorRepository) GetHistory(deviceID int, duration string) ([]model.SensorData, error) {

	var data []model.SensorData

	var timeFilter string

	switch duration {
	case "1h":
		timeFilter = "DATEADD(HOUR, -1, GETDATE())"
	case "1d":
		timeFilter = "DATEADD(DAY, -1, GETDATE())"
	case "1w":
		timeFilter = "DATEADD(WEEK, -1, GETDATE())"
	default:
		timeFilter = "DATEADD(DAY, -1, GETDATE())"
	}

	err := r.DB.
		Where("device_id = ? AND created_at >= "+timeFilter, deviceID).
		Order("created_at ASC").
		Find(&data).Error

	return data, err
}
