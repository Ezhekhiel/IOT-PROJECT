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
