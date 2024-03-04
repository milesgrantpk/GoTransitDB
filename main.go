// main.go
package main

import (
	"milesgrant/gotransitdb/controllers"
	"milesgrant/gotransitdb/models"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    models.ConnectDatabase()

    router.POST("/stations", controllers.CreateStation)
    router.GET("/stations", controllers.GetAllStations)
    router.GET("/stations/:id", controllers.GetStationById)
    router.PATCH("/stations/:id", controllers.UpdateStation)
    router.DELETE("/stations/:id", controllers.DeleteStation)

    router.Run("localhost:8080")
}