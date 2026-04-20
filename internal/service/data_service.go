package service

import (
	"errors"
	"fmt"

	"github.com/Ezhekhiel/IOT-PROJECT/internal/model"
	"github.com/Ezhekhiel/IOT-PROJECT/internal/repository"
)

type DataService struct {
	StandarRepo     repository.StandardRepository
	CellRunningRepo repository.CellRunningModelRepository
}
type CellRequest struct {
	Name          string `json:"name"`
	LocationId    *int   `json:"location_id"`
	TargetPerHour *int   `json:"target_per_hour"`
}
type LocationRequest struct {
	Name string `json:"name"`
}
type ModelRequest struct {
	Name string `json:"name"`
}
type ProcessRequest struct {
	Name string `json:"name"`
}

func (s *DataService) GetModelProcessStandards() (interface{}, error) {

	data, err := s.StandarRepo.GetModelProcessStandards()
	if err != nil {
		return nil, err
	}

	// format untuk chart
	var result []map[string]interface{}
	for _, d := range data {
		result = append(result, map[string]interface{}{
			"id":              d.ID,
			"model_id":        d.ModelId,
			"process_id":      d.ProcessId,
			"cell_id":         d.CellId,
			"name":            d.Name,
			"target_per_hour": d.TargetPerHour,
			"Model":           d.Model,
			"Process":         d.Process,
			"MinPressure":     d.MinPressure,
			"MaxPressure":     d.MaxPressure,
			"MinTemperature":  d.MinTemperature,
			"MaxTemperature":  d.MaxTemperature,
		})
	}
	return result, nil
}
func (s *DataService) GetAllCell() (interface{}, error) {

	data, err := s.CellRunningRepo.GetAllCell()
	if err != nil {
		return nil, err
	}

	// format untuk chart
	var result []map[string]interface{}
	for _, d := range data {
		result = append(result, map[string]interface{}{
			"id":              d.ID,
			"name":            d.Name,
			"location_id":     d.LocationId,
			"location":        d.Location,
			"target_per_hour": d.TargetPerHour,
		})
	}
	return result, nil
}
func (s *DataService) AddCell(data CellRequest) (string, error) {

	// 1. Find Device
	_, err := s.GetAllLocation(*data.LocationId)
	if err != nil {
		return "", errors.New("location not found")
	}

	status := "OK"

	// 5. SAVE SENSOR
	sensor := model.Cell{
		Name:          data.Name,
		LocationId:    *data.LocationId,
		TargetPerHour: *data.TargetPerHour,
	}

	err = s.CellRunningRepo.SaveCell(sensor)
	if err != nil {
		return "", err
	}

	return status, nil
}
func (s *DataService) GetAllLocation(locationId int) (interface{}, error) {

	data, err := s.CellRunningRepo.GetAllLocation(locationId)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}

	for _, d := range data {
		result = append(result, map[string]interface{}{
			"id":   d.ID,
			"name": d.Name,
		})
	}

	return result, nil
}
func (s *DataService) AddLocation(data LocationRequest) (string, error) {

	exists, err := s.CellRunningRepo.IsLocationExists(data.Name)
	if err != nil {
		return "", err
	}

	if exists {
		return "", fmt.Errorf("Location already exists")
	}

	location := model.Location{
		Name: data.Name,
	}

	err = s.CellRunningRepo.SaveLocation(location)
	if err != nil {
		return "", err
	}

	return "OK", nil
}
func (s *DataService) GetAllModel() (interface{}, error) {

	data, err := s.CellRunningRepo.GetAllModel()
	if err != nil {
		return nil, err
	}

	// format untuk chart
	var result []map[string]interface{}
	for _, d := range data {
		result = append(result, map[string]interface{}{
			"id":   d.ID,
			"name": d.Name,
		})
	}
	return result, nil
}
func (s *DataService) AddModel(data ModelRequest) (string, error) {

	exists, err := s.CellRunningRepo.IsModelExist(data.Name)
	if err != nil {
		return "", err
	}

	if exists {
		return "", fmt.Errorf("Model already exists")
	}

	model := model.Model{
		Name: data.Name,
	}

	err = s.CellRunningRepo.SaveModel(model)
	if err != nil {
		return "", err
	}

	return "OK", nil
}
func (s *DataService) AddProcess(data ProcessRequest) (string, error) {
	exist, err := s.CellRunningRepo.IsProcessExist(data.Name)
	if err != nil {
		return "", err
	}
	if exist {
		return "", fmt.Errorf("Process already exist")
	}
	process := model.Process{
		Name: data.Name,
	}
	err = s.CellRunningRepo.SaveProcess(process)
	if err != nil {
		return "", err
	}
	return "OK", nil
}
func (s *DataService) GetAllProcess() (interface{}, error) {

	data, err := s.CellRunningRepo.GetAllProcess()
	if err != nil {
		return nil, err
	}

	// format untuk chart
	var result []map[string]interface{}
	for _, d := range data {
		result = append(result, map[string]interface{}{
			"id":   d.ID,
			"name": d.Name,
		})
	}
	return result, nil
}
