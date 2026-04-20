package repository

import (
	"time"

	"github.com/Ezhekhiel/IOT-PROJECT/internal/model"

	"gorm.io/gorm"
)

type SensorRepository struct {
	DB *gorm.DB
}
type HistoryAggregated struct {
	TimeGroup   time.Time `gorm:"column:time_group"`
	Pressure    *float64  `gorm:"column:pressure"`
	Temperature *float64  `gorm:"column:temperature"`
	Timer       *float64  `gorm:"column:timer"`
	DeviceCode  string    `gorm:"column:device_code"`
}
type TimeConfig struct {
	GroupFormat string
	RangeFilter string
}

func (r *SensorRepository) Save(data model.SensorData) error {
	return r.DB.Create(&data).Error
}
func (r *SensorRepository) GetLatestByDevice(deviceID int) (model.SensorData, error) {
	var data model.SensorData
	err := r.DB.
		Where("device_id = ?", deviceID).
		Order("created_at DESC").
		First(&data).Error

	return data, err
}
func (r *SensorRepository) GetLatestAllDevices() ([]model.SensorData, error) {

	var results []model.SensorData

	query := `
		SELECT *
		FROM (
			SELECT *,
				   ROW_NUMBER() OVER (
					   PARTITION BY device_id 
					   ORDER BY created_at DESC
				   ) as rn
			FROM sensor_data
		) t
		WHERE rn = 1
	`

	err := r.DB.Raw(query).Scan(&results).Error

	return results, err
}
func (r *SensorRepository) GetHistory(deviceID int, duration string) ([]model.SensorData, error) {

	var data []model.SensorData

	var timeFilter string

	switch duration {
	case "1h":
		timeFilter = "DATEADD(HOUR, -1, GETDATE())"
	case "1d":
		timeFilter = "DATEADD(DAY, -1, GETDATE())"
	case "1w":
		timeFilter = "DATEADD(WEEK, -1, GETDATE())"
	default:
		timeFilter = "DATEADD(DAY, -1, GETDATE())"
	}

	err := r.DB.
		Where("device_id = ? AND created_at >= "+timeFilter, deviceID).
		Order("created_at ASC").
		Find(&data).Error

	return data, err
}
func (r *SensorRepository) GetHistoryAggregated(duration string) ([]HistoryAggregated, error) {

	var results []HistoryAggregated

	// 🔥 CONFIG MAP
	configs := map[string]TimeConfig{
		"1h": {
			GroupFormat: "DATEADD(MINUTE, DATEDIFF(MINUTE, 0, a.created_at), 0)",
			RangeFilter: "DATEADD(HOUR, -1, GETDATE())",
		},
		"1d": {
			GroupFormat: "DATEADD(HOUR, DATEDIFF(HOUR, 0, a.created_at), 0)",
			RangeFilter: "DATEADD(DAY, -1, GETDATE())",
		},
		"1w": {
			GroupFormat: "CAST(a.created_at AS DATE)",
			RangeFilter: "DATEADD(WEEK, -1, GETDATE())",
		},
	}

	cfg, ok := configs[duration]
	if !ok {
		cfg = configs["1d"] // default
	}

	// 🔥 dynamic query
	query := `
		SELECT 
			` + cfg.GroupFormat + ` as time_group,
			AVG(a.pressure) as pressure,
			AVG(a.temperature) as temperature,
			AVG(a.timer) as timer,
			b.device_code

		FROM sensor_data as a
		JOIN devices as b ON a.device_id = b.id

		WHERE a.created_at >= ` + cfg.RangeFilter + `

		GROUP BY 
			` + cfg.GroupFormat + `,
			b.device_code

		ORDER BY b.device_code, time_group ASC
		`

	err := r.DB.Raw(query).Scan(&results).Error

	return results, err
}
