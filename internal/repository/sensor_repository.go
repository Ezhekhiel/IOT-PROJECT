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
