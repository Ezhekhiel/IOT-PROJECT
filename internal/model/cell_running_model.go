package model

import "time"

type CellRunningModel struct {
	ID        int
	CellID    int
	ModelID   int
	StartTime time.Time
}

//aktivkan ini kalo mau setting manual table
// func (Device) TableName() string {
// 	return "cell_running_models"
// }
