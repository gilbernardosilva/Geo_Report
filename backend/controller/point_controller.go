package controller

import (
	"geo_report_api/dto"
	"geo_report_api/service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePoint creates a new point associated with an area.
//
//	@Summary		Create Point
//	@Description	Creates a new point associated with an area.
//	@Tags			Point
//	@Accept			json
//	@Produce		json
//	@Param			area_id	path		uint64			true	"Area ID"
//	@Param			request	body		dto.PointDTO	true	"Point data"
//	@Success		200		{object}	dto.PointResponseDTO
//	@Failure		400		{object}	dto.ErrorResponse
//	@Router			/area/{area_id}/point [post]
func CreatePoint(c *gin.Context) {
	areaID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid area ID",
			"error":   err.Error(),
		})
		return
	}

	var pointDTO dto.PointDTO
	if err := c.ShouldBindJSON(&pointDTO); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	createdPoint, err := service.CreatePoint(pointDTO, areaID)
	if err != nil {
		log.Printf("Error creating point: %v", err)
		c.JSON(400, gin.H{
			"message": "Failed to create point",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Point created successfully",
		"point":   createdPoint,
	})
}

// GetPointsByAreaID retrieves all points associated with an area.
//
//	@Summary		Get Points by Area ID
//	@Description	Retrieves all points associated with an area by its ID.
//	@Tags			Point
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64	true	"Area ID"
//	@Success		200		{array}		dto.PointResponseDTO
//	@Failure		400		{object}	dto.ErrorResponse
//	@Router			/area/{id}/point [get]
func GetPointsByAreaID(c *gin.Context) {
	areaID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid area ID",
			"error":   err.Error(),
		})
		return
	}

	points, err := service.GetPointsByAreaID(areaID)
	if err != nil {
		log.Printf("Error retrieving points: %v", err)
		c.JSON(400, gin.H{
			"message": "Failed to retrieve points",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Points retrieved successfully",
		"points":  points,
	})
}

// UpdatePoint updates a point associated with an area by its ID.
//
//	@Summary		Update Point
//	@Description	Updates a point associated with an area by its ID.
//	@Tags			Point
//	@Accept			json
//	@Produce		json
//	@Param			id		path		uint64				true	"Point ID"
//	@Param			request	body		dto.PointUpdateDTO	true	"Updated point data"
//	@Success		200		{object}	dto.PointResponseDTO
//	@Failure		400		{object}	dto.ErrorResponse
//	@Router			/point/{id} [put]
func UpdatePoint(c *gin.Context) {
	pointID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid point ID",
			"error":   err.Error(),
		})
		return
	}

	var updateDTO dto.PointUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	updatedPoint, err := service.UpdatePoint(pointID, updateDTO)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to update point",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Point updated successfully",
		"point":   updatedPoint,
	})
}

// DeletePoint deletes a point associated with an area by its ID.
//
//	@Summary		Delete Point
//	@Description	Deletes a point associated with an area by its ID.
//	@Tags			Point
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64	true	"Point ID"
//	@Success		200	{object}	dto.SuccessResponse
//	@Failure		400	{object}	dto.ErrorResponse
//	@Router			/point/{id} [delete]
func DeletePoint(c *gin.Context) {
	pointID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid point ID",
			"error":   err.Error(),
		})
		return
	}

	err = service.DeletePoint(pointID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed to delete point",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Point deleted successfully",
	})
}
