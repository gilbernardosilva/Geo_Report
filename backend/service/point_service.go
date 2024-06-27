package service

import (
	"errors"
	"geo_report_api/dto"
	"geo_report_api/model"
	"geo_report_api/repository"
	"log"

	"github.com/mashingan/smapping"
)

func CreatePoint(pointDTO dto.PointDTO, areaID uint64) (dto.PointResponseDTO, error) {
	var point model.Point
	err := smapping.FillStruct(&point, smapping.MapFields(&pointDTO))
	if err != nil {
		log.Printf("Failed to map point DTO to point struct: %v", err)
		return dto.PointResponseDTO{}, errors.New("failed to map point")
	}

	area, err := repository.GetAreaByID(areaID)
	if err != nil {
		log.Printf("Failed to get area by area ID: %v", err)
		return dto.PointResponseDTO{}, errors.New("failed to get area")
	}

	point.Area = area
	point.AreaID = areaID

	err = repository.CreatePoint(&point)
	if err != nil {
		return dto.PointResponseDTO{}, err
	}
	return pointToDTO(point), nil
}

func GetPointsByAreaID(areaID uint64) ([]dto.PointResponseDTO, error) {
	points, err := repository.GetPointsByAreaID(areaID)
	if err != nil {
		return nil, err
	}
	var dtos []dto.PointResponseDTO
	for _, point := range points {
		dtos = append(dtos, pointToDTO(point))
	}
	return dtos, nil
}

func UpdatePoint(pointID uint64, update dto.PointUpdateDTO) (dto.PointResponseDTO, error) {
	modelPoint, err := repository.GetPointByID(pointID)
	if err != nil {
		return dto.PointResponseDTO{}, err
	}

	modelPoint.Latitude = update.Latitude
	modelPoint.Longitude = update.Longitude

	modelPoint, err = repository.UpdatePoint(modelPoint)
	if err != nil {
		return dto.PointResponseDTO{}, err
	}

	return pointToDTO(modelPoint), nil
}

func DeletePoint(pointID uint64) error {
	return repository.DeletePoint(pointID)
}

func pointToDTO(point model.Point) dto.PointResponseDTO {
	return dto.PointResponseDTO{
		ID:        point.ID,
		Latitude:  point.Latitude,
		Longitude: point.Longitude,
	}
}
