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
