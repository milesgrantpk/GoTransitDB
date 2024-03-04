// controllers/station.go
package controllers

import (
    "net/http"
    "milesgrant/gotransitdb/models"

    "github.com/gin-gonic/gin"
)

type CreateStationInput struct {
    Name    	string 	`json:"name" binding:"required"`
    Line    	string 	`json:"line" binding:"required"`
	X	    	*int64 	`json:"x" binding:"required"`
	Y	    	*int64 	`json:"y" binding:"required"`
}

func CreateStation(c *gin.Context) {
    var input CreateStationInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    station := models.Station{NAME: input.Name, LINE: input.Line, X: input.X, Y: input.Y}
    models.DB.Create(&station)

    c.JSON(http.StatusOK, gin.H{"data": station})
}

func GetAllStations(c *gin.Context) {
    var stations []models.Station
    models.DB.Find(&stations)

    c.JSON(http.StatusOK, gin.H{"data": stations})
}

func GetStationById(c *gin.Context) {
    var station models.Station

    if err := models.DB.Where("id = ?", c.Param("id")).First(&station).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": station})
}

type UpdateStationInput struct {
    Name    	string 	`json:"name"`
    Line    	string 	`json:"line"`
	X	    	*int64 	`json:"x"`
	Y	    	*int64 	`json:"y"`
}

func UpdateStation(c *gin.Context) {
    var station models.Station
    if err := models.DB.Where("id = ?", c.Param("id")).First(&station).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
        return
    }

    var input UpdateStationInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedStation := models.Station{NAME: input.Name, LINE: input.Line, X: input.X, Y: input.Y}

    models.DB.Model(&station).Updates(&updatedStation)
    c.JSON(http.StatusOK, gin.H{"data": station})
}

func DeleteStation(c *gin.Context) {
    var station models.Station
    if err := models.DB.Where("id = ?", c.Param("id")).First(&station).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
        return
    }

    models.DB.Delete(&station)
    c.JSON(http.StatusOK, gin.H{"data": "success"})
}
