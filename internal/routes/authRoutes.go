package routes

import (
    "pumplepet-server/internal/controller/auth"
    "github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
    authRoutes := router.Group("/auth")
    {
        authRoutes.POST("/register", auth.Register)
        authRoutes.POST("/login", auth.Login)
    }
}