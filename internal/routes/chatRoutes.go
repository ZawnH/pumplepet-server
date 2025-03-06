package routes

import (
	"pumplepet-server/internal/controller/chat"
	"pumplepet-server/internal/middleware"
	"pumplepet-server/internal/websocket"

	"github.com/gin-gonic/gin"
)

func ChatRoutes(router *gin.Engine, manager *websocket.Manager) {
	chatGroup := router.Group("/chat")
	chatGroup.Use(middleware.AuthMiddleware())
	chatGroup.Use(func(c *gin.Context) {
		c.Set("ws_manager", manager)
		c.Next()
	})
	{
		chatGroup.GET("/ws", chat.HandleWebSocket(manager))
		chatGroup.GET("/history", chat.GetChatHistory)
		chatGroup.POST("/send", chat.SendMessage)
		chatGroup.GET("/messages/:user_id", chat.GetChatWithUser)
	}
}
