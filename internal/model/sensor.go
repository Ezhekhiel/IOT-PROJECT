package model

import "time"

type SensorData struct {
	ID          int64     `gorm:"column:id;primaryKey"`
	DeviceID    int       `gorm:"column:device_id"`
	Pressure    *float64  `gorm:"column:pressure"`
	Temperature *float64  `gorm:"column:temperature"`
	Timer       *int      `gorm:"column:timer"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}
