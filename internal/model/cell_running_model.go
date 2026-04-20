package model

import "time"

type CellRunningModel struct {
	ID        int
	CellID    int
	ModelID   int
	StartTime time.Time
}
type Cell struct {
	ID            int    `gorm:"column:id;primaryKey"`
	Name          string `gorm:"column:name"`
	LocationId    int    `gorm:"column:location_id"`
	TargetPerHour int    `gorm:"column:target_per_hour"`
}
type CellDatabase struct {
	ID            int    `gorm:"column:id;primaryKey"`
	Name          string `gorm:"column:name"`
	LocationId    int    `gorm:"column:location_id"`
	Location      string `gorm:"column:location"`
	TargetPerHour int    `gorm:"column:target_per_hour"`
}
type Location struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}
type Model struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}
type Process struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}

//aktivkan ini kalo mau setting manual table
// func (Device) TableName() string {
// 	return "cell_running_models"
// }
