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
func (r *CellRunningModelRepository) GetAllCell() ([]model.CellDatabase, error) {
	var result []model.CellDatabase

	query := `
		select a.*, b.name as location
		from cells as a 
		join locations as b on a.location_id = b.id 
	`

	err := r.DB.Raw(query).Scan(&result).Error
	return result, err
}
func (r *CellRunningModelRepository) SaveCell(data model.Cell) error {
	return r.DB.Create(&data).Error
}
func (r *CellRunningModelRepository) GetAllLocation(locationId int) ([]model.Location, error) {
	var result []model.Location

	db := r.DB

	if locationId != 0 {
		db = db.Where("id = ?", locationId)
	}

	err := db.Find(&result).Error

	return result, err
}
func (r *CellRunningModelRepository) IsLocationExists(name string) (bool, error) {
	var count int64

	err := r.DB.
		Model(&model.Location{}).
		Where("name = ?", name).
		Count(&count).Error

	return count > 0, err
}
func (r *CellRunningModelRepository) IsModelExist(name string) (bool, error) {
	var count int64

	err := r.DB.
		Model(&model.Model{}).
		Where("name = ?", name).
		Count(&count).Error

	return count > 0, err
}
func (r *CellRunningModelRepository) IsProcessExist(name string) (bool, error) {
	var count int64

	err := r.DB.
		Model(&model.Process{}).
		Where("name = ?", name).
		Count(&count).Error

	return count > 0, err
}
func (r *CellRunningModelRepository) SaveLocation(data model.Location) error {
	return r.DB.Create(&data).Error
}
func (r *CellRunningModelRepository) GetAllModel() ([]model.Model, error) {
	var result []model.Model

	err := r.DB.
		Find(&result).Error

	return result, err
}
func (r *CellRunningModelRepository) SaveModel(data model.Model) error {
	return r.DB.Create(&data).Error
}
func (r *CellRunningModelRepository) GetAllProcess() ([]model.Process, error) {
	var result []model.Process

	err := r.DB.
		Find(&result).Error
	if err != nil {

	}

	return result, err
}
func (r *CellRunningModelRepository) SaveProcess(data model.Process) error {
	return r.DB.Create(&data).Error
}
