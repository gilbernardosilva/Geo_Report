package repository

import (
	"errors"
	"geo_report_api/config"
	"geo_report_api/model"
)

func GetAllAreas() ([]model.Area, error) {
	var areas []model.Area
	err := config.Db.Find(&areas).Error
	return areas, err
}

func CreateArea(area model.Area) error {
	return config.Db.Create(&area).Error
}

func GetAreaByID(areaID uint64) (model.Area, error) {
	var area model.Area
	config.Db.First(&area, areaID)
	if area.ID != 0 {
		return area, nil
	}
	return area, errors.New("area not found")
}

func UpdateArea(area model.Area) (model.Area, error) {
	if _, err := GetAreaByID(area.ID); err == nil {
		config.Db.Save(&area)
		config.Db.Find(&area)
		return area, nil
	}
	return model.Area{}, errors.New("error updating area")
}

func DeleteArea(areaID uint64) error {
	return config.Db.Delete(&model.Area{}, areaID).Error
}
