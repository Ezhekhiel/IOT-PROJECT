package model

import "time"

type SensorData struct {
	ID          int64
	DeviceID    int
	Pressure    *float64
	Temperature *float64
	Timer       *int
	CreatedAt   time.Time
}
