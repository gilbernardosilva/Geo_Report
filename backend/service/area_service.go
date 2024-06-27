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

	var areaDTOs []dto.AreaResponseDTO
	for _, area := range areas {
		var pointDTOs []dto.PointResponseDTO
		for _, point := range area.Points {
			pointDTO := dto.PointResponseDTO{
				ID:        point.ID,
				Latitude:  point.Latitude,
				Longitude: point.Longitude,
			}
			pointDTOs = append(pointDTOs, pointDTO)
		}

		areaDTO := dto.AreaResponseDTO{
			ID:     area.ID,
			Name:   area.Name,
			Points: pointDTOs,
		}
		areaDTOs = append(areaDTOs, areaDTO)
	}
	return areaDTOs, nil
}

// creates an area without points... there's another endpoint to create points
func CreateArea(areaDTO dto.AreaCreateDTO) error {
	var area model.Area
	err := smapping.FillStruct(&area, smapping.MapFields(&areaDTO))
	if err != nil {
		log.Printf("Failed to map area DTO to area struct: %v", err)
		return errors.New("failed to map area")
	}

	// points will just be an empty array since they'll be added later on
	area.Points = []model.Point{}

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
