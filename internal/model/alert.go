package model

import "time"

type Alert struct {
	ID         int64
	DeviceID   int
	Message    string
	CreatedAt  time.Time
	ResolvedAt *time.Time
}
