package model

import "time"

type Alert struct {
	ID        int64 `gorm:"primaryKey"`
	DeviceID  int
	Message   string
	Status    string // "ACTIVE" | "RESOLVED"
	CreatedAt time.Time
	UpdatedAt time.Time
}

//aktivkan ini kalo mau setting manual table
// func (Device) TableName() string {
// 	return "alerts"
// }
