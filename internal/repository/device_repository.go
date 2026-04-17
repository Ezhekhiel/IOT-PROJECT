package repository

import (
	"github.com/Ezhekhiel/IOT-PROJECT/internal/model"
	"gorm.io/gorm"
)

type DeviceRepository struct {
	DB *gorm.DB
}

func (r *DeviceRepository) FindByCode(code string) (model.Device, error) {
	var device model.Device
	err := r.DB.Where("device_code = ?", code).First(&device).Error
	return device, err
}
