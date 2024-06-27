package controller

import (
	"geo_report_api/dto"
	"geo_report_api/model"
	"geo_report_api/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

// GetAllAreas retrieves all areas.
//
//	@Summary	Get all areas
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}		dto.AreaResponseDTO
//	@Failure	500	{object}	dto.ErrorResponse
//	@Router		/area [get]
func GetAllAreas(c *gin.Context) {
	areas, err := service.GetAllAreas()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to fetch areas",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "areas successfully retrieved",
		"areas":   areas,
	})
}

// CreateArea handles the creation of a new area.
//
//	@Summary	Create a new area
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.AreaCreateDTO	true	"Area data"
//	@Success	200		{object}	dto.AreaResponseDTO
//	@Failure	400		{object}	dto.ErrorResponse
//	@Router		/area [post]
func CreateArea(c *gin.Context) {
	var area dto.AreaCreateDTO
	if err := c.ShouldBindJSON(&area); err != nil {
		c.JSON(400, gin.H{

			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	err := service.CreateArea(area)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create area",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "area successfully created",
	})
}

// GetAreaByID retrieves an area by ID.
//
//	@Summary	Get an area by ID
//	@Accept		json
//	@Produce	json
//	@Param		id	path		uint64	true	"Area ID"
//	@Success	200	{object}	dto.AreaResponseDTO
//	@Failure	400	{object}	dto.ErrorResponse
//	@Router		/area/{id} [get]
func GetAreaByID(c *gin.Context) {
	areaID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid area ID",
		})
		return
	}

	area, err := service.GetAreaByID(areaID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Area not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "area successfully retrieved",
		"area":    area,
	})
}

// UpdateArea updates an existing area.
//
//	@Summary	Update an area
//	@Accept		json
//	@Produce	json
//	@Param		id		path		uint64				true	"Area ID"
//	@Param		request	body		dto.AreaUpdateDTO	true	"Area data"
//	@Success	200		{object}	dto.AreaResponseDTO
//	@Failure	400		{object}	dto.ErrorResponse
//	@Router		/area/{id} [put]
func UpdateArea(c *gin.Context) {
	var area model.Area
	areaID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid area ID",
		})
		return
	}

	var areaDTO dto.AreaUpdateDTO
	if err := c.ShouldBindJSON(&areaDTO); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
		return
	}

	err = smapping.FillStruct(&area, smapping.MapFields(&areaDTO))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "error mapping struct",
		})
		return
	}
	area.ID = areaID
	areaResponse, err := service.EditArea(area)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Area not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "area successfully updated",
		"area":    areaResponse,
	})
}

// DeleteArea deletes an area by ID.
//
//	@Summary	Delete an area
//	@Accept		json
//	@Produce	json
//	@Param		id	path		uint64	true	"Area ID"
//	@Success	200	{object}	dto.SuccessResponse
//	@Failure	400	{object}	dto.ErrorResponse
//	@Router		/area/{id} [delete]
func DeleteArea(c *gin.Context) {
	areaID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid area ID",
		})
		return
	}

	err = service.DeleteArea(areaID)
	if err != nil {

		c.JSON(400, gin.H{
			"error": "Can't delete area because it's referenced by an authority",
		})
		return

	}

	c.JSON(200, gin.H{
		"message": "Area deleted successfully",
	})
}
