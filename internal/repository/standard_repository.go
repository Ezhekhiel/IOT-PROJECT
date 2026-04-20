package repository

import (
	"github.com/Ezhekhiel/IOT-PROJECT/internal/model"
	"gorm.io/gorm"
)

type StandardRepository struct {
	DB *gorm.DB
}

func (r *StandardRepository) Find(modelID int, processID int) (model.ModelProcessStandard, error) {
	var standard model.ModelProcessStandard

	err := r.DB.
		Where("model_id = ? AND process_id = ?", modelID, processID).
		First(&standard).Error

	return standard, err
}
func (r *StandardRepository) GetModelProcessStandards() ([]model.ModelProcessStandardFull, error) {
	var result []model.ModelProcessStandardFull

	query := `
		select a.*, b.name, b.target_per_hour, c.name as model, d.name as process 
		from model_process_standards as a 
		join cells as b on a.cell_id = b.id 
		join models as c on a.model_id = c.id 
		join processes as d on a.process_id = d.id
	`

	err := r.DB.Raw(query).Scan(&result).Error
	return result, err
}
