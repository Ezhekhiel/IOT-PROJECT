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
