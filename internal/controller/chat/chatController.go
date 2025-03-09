package chat

import (
	"net/http"
	"pumplepet-server/internal/middleware"
	"pumplepet-server/internal/service/chat"
	"pumplepet-server/internal/websocket"

	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	gorilla "github.com/gorilla/websocket"
)

var upgrader = gorilla.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: implement proper origin check in prod
	},
}

func HandleWebSocket(manager *websocket.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from query parameter
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			return
		}

		// Validate token and get user ID
		userID, err := middleware.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Upgrade connection
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not upgrade connection"})
			return
		}

		// Create client with validated user ID
		client := &websocket.Client{
			Conn:   conn,
			Send:   make(chan []byte),
			UserID: userID,
		}

		manager.RegisterClient(client)

		go client.WritePump()
		go client.ReadPump(manager, chat.SaveMessage)
	}
}

func GetChatHistory(c *gin.Context) {
	messages, err := chat.GetMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching messages"})
		return
	}
	c.JSON(http.StatusOK, messages)
}

func SendMessage(c *gin.Context) {
	var payload MessagePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id") // Get user ID from context
	message, err := chat.SaveMessage(payload.Content, userID.(uint), payload.RecipientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving message"})
		return
	}

	// Get the WebSocket manager from the application context
	wsManager := c.MustGet("ws_manager").(*websocket.Manager)

	// Broadcast the message to all connected clients
	if err := chat.BroadcastMessage(wsManager, message); err != nil {
		log.Println("Error broadcasting message:", err)
	}

	c.JSON(http.StatusOK, message)
}

func GetChatWithUser(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	otherUserID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	messages, err := chat.GetMessagesBetweenUsers(userID, uint(otherUserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}
