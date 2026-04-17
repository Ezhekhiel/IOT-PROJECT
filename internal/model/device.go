package model

type Device struct {
	ID         int    `gorm:"column:id;primaryKey"`
	DeviceCode string `gorm:"column:device_code"`
	CellID     int    `gorm:"column:cell_id"`
	ProcessID  int    `gorm:"column:process_id"`
}
