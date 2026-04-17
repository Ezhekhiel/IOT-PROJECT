package repository

import (
	"time"

	"github.com/Ezhekhiel/IOT-PROJECT/internal/model"
	"gorm.io/gorm"
)

type AlertRepository struct {
	DB *gorm.DB
}

func (r *AlertRepository) Create(alert model.Alert) error {
	return r.DB.Create(&alert).Error
}
func (r *AlertRepository) GetActive() ([]model.Alert, error) {
	var alerts []model.Alert

	err := r.DB.
		Where("status = ?", "ACTIVE").
		Order("created_at DESC").
		Find(&alerts).Error

	return alerts, err
}
func (r *AlertRepository) ResolveByDevice(deviceID int) error {
	return r.DB.
		Model(&model.Alert{}).
		Where("device_id = ? AND status = ?", deviceID, "ACTIVE").
		Updates(map[string]interface{}{
			"status": "RESOLVED",
		}).Error
}
func (r *AlertRepository) ResolveActiveAlert(deviceID int) {
	now := time.Now()

	r.DB.Model(&model.Alert{}).
		Where("device_id = ? AND resolved_at IS NULL", deviceID).
		Update("resolved_at", now)
}
