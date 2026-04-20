package service

import "github.com/Ezhekhiel/IOT-PROJECT/internal/repository"

type DataService struct {
	DataRepo repository.StandardRepository
}

func (s *DataService) GetModelProcessStandards() (interface{}, error) {

	data, err := s.DataRepo.GetModelProcessStandards()
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
