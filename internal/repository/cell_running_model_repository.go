package repository

import (
	"github.com/Ezhekhiel/IOT-PROJECT/internal/model"
	"gorm.io/gorm"
)

type CellRunningModelRepository struct {
	DB *gorm.DB
}

func (r *CellRunningModelRepository) FindLatestByCell(cellID int) (model.CellRunningModel, error) {
	var result model.CellRunningModel
	err := r.DB.
		Where("cell_id = ?", cellID).
		Order("start_time DESC").
		First(&result).Error

	return result, err
}
