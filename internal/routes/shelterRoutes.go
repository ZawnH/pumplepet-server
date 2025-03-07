package routes

import (
	"pumplepet-server/internal/controller/shelter"

	"github.com/gin-gonic/gin"
)

func ShelterRoutes(router *gin.Engine) {
	shelterRoutes := router.Group("/shelters")
	{
		shelterRoutes.GET("", shelter.GetShelters)
		shelterRoutes.GET("/:id", shelter.GetShelterById)
	}
}
