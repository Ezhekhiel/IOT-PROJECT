package model

type ModelProcessStandard struct {
	ID             int
	ModelID        int
	ProcessID      int
	MinPressure    *float64
	MaxPressure    *float64
	MinTemperature *float64
	MaxTemperature *float64
	MaxTimer       *int
}
type ModelProcessStandardFull struct {
	ID             int    `gorm:"column:id;primaryKey"`
	ModelId        int    `gorm:"column:model_id"`
	ProcessId      int    `gorm:"column:process_id"`
	CellId         int    `gorm:"column:cell_id"`
	Name           string `gorm:"column:name"`
	TargetPerHour  int    `gorm:"column:target_per_hour"`
	Model          string `gorm:"column:model"`
	Process        string `gorm:"column:process"`
	MinPressure    int    `gorm:"column:min_pressure"`
	MaxPressure    int    `gorm:"column:max_pressure"`
	MinTemperature int    `gorm:"column:min_temperature"`
	MaxTemperature int    `gorm:"column:max_temperature"`
}

//aktivkan ini kalo mau setting manual table
// func (Device) TableName() string {
// 	return "model_process_standards"
// }
