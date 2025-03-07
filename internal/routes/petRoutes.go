package routes

import (
	"pumplepet-server/internal/controller/pet"

	"github.com/gin-gonic/gin"
)

func PetRoutes(router *gin.Engine) {
	petRoutes := router.Group("/pets")
	{
		petRoutes.GET("", pet.GetPets)
		petRoutes.GET("/:id", pet.GetPetById)
		petRoutes.POST("/:id/like", pet.LikePet)
		petRoutes.POST("/:id/dislike", pet.DislikePet)
	}
}
