package repository

import (
	"errors"
	"geo_report_api/config"
	"geo_report_api/model"
)

func CreatePoint(point *model.Point) error {
	if err := config.Db.Preload("Area").Create(point).Error; err != nil {
		return err
	}
	return nil
}

func GetPointsByAreaID(areaID uint64) ([]model.Point, error) {
	var points []model.Point
	err := config.Db.Where("area_id = ?", areaID).Find(&points).Error
	return points, err
}

func GetPointByID(pointID uint64) (model.Point, error) {
	var point model.Point
	config.Db.First(&point, pointID)
	if point.ID != 0 {
		return point, nil
	}
	return point, errors.New("point not found")
}

func UpdatePoint(point model.Point) (model.Point, error) {
	if _, err := GetPointByID(point.ID); err == nil {
		config.Db.Save(&point)
		config.Db.Find(&point)
		return point, nil
	}
	return model.Point{}, errors.New("error updating point")
}

func DeletePoint(pointID uint64) error {
	return config.Db.Delete(&model.Point{}, pointID).Error
}
