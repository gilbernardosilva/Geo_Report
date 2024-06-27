package dto

import "time"

type PointDTO struct {
	Latitude  string `json:"latitude" binding:"required"`
	Longitude string `json:"longitude" binding:"required"`
}

type PointUpdateDTO struct {
	AreaID    uint64 `json:"area_id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type PointResponseDTO struct {
	ID        uint64    `json:"id"`
	AreaID    uint64    `json:"area_id"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PointResponseForAreaDTO struct {
	ID        uint64 `json:"id"`
	AreaID    uint64 `json:"area_id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
