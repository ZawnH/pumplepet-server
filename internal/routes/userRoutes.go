package routes

import (
	"pumplepet-server/internal/controller/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", user.GetUsers)
		userRoutes.GET("/:id", user.GetUserById)
	}
}
