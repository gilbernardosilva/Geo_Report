package dto

type AreaCreateDTO struct {
	Name string `json:"name" binding:"required"`
}

type AreaUpdateDTO struct {
	Name string `json:"name"`
}

type AreaResponseDTO struct {
	ID     uint64             `json:"id"`
	Name   string             `json:"name"`
	Points []PointResponseDTO `json:"points"`
}
