package service

import (
	"errors"
	"geo_report_api/dto"
	"geo_report_api/model"
	"geo_report_api/repository"
	"log"

	"github.com/mashingan/smapping"
)

// returns all areas with the respective points
func GetAllAreas() ([]dto.AreaResponseDTO, error) {
	areas, err := repository.GetAllAreas()
	if err != nil {
		return nil, err
	}

	areaDTOs := make([]dto.AreaResponseDTO, len(areas))
	for i, area := range areas {
		areaDTOs[i] = dto.AreaResponseDTO{
			ID:        area.ID,
			Name:      area.Name,
			Latitude:  area.Latitude,
			Longitude: area.Longitude,
			Radius:    area.Radius,
		}
	}

	return areaDTOs, nil
}

func CreateArea(areaDTO dto.AreaCreateDTO) error {
	var area model.Area

	err := smapping.FillStruct(&area, smapping.MapFields(&areaDTO))
	if err != nil {
		log.Printf("Failed to map area DTO to area struct: %v", err)
		return errors.New("failed to map area")
	}

	return repository.CreateArea(area)
}

func GetAreaByID(areaID uint64) (model.Area, error) {
	return repository.GetAreaByID(areaID)
}

func EditArea(area model.Area) (model.Area, error) {
	if updatedArea, err := repository.UpdateArea(area); err == nil {
		return updatedArea, nil
	}
	return model.Area{}, errors.New("error editing area")
}

func DeleteArea(areaID uint64) error {
	return repository.DeleteArea(areaID)
}
