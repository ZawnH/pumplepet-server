package websocket

import (
	"encoding/json"
	"log"
	"pumplepet-server/internal/model"

	"github.com/gorilla/websocket"
)

func (c *Client) WritePump() {
	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error writing message:", err)
			return
		}
	}
	// If the channel is closed, close the connection
	c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
	c.Conn.Close()
}

func (c *Client) ReadPump(manager *Manager, saveMessageFunc func(content string, senderID uint, recipientID uint) (*model.Message, error)) {
	defer func() {
		manager.unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Parse the message
		var messageData struct {
			Content     string `json:"content"`
			RecipientID uint   `json:"recipient_id"`
		}
		if err := json.Unmarshal(message, &messageData); err != nil {
			log.Println("Error parsing message:", err)
			continue
		}

		if messageData.Content == "" || messageData.RecipientID == 0 {
			log.Println("Invalid message format: missing content or recipient_id")
			continue
		}

		// Use the callback instead of directly calling chat.SaveMessage
		savedMessage, err := saveMessageFunc(messageData.Content, c.UserID, messageData.RecipientID)
		if err != nil {
			log.Println("Error saving message:", err)
			continue
		}

		// Marshal the saved message (which includes user data)
		messageJSON, err := json.Marshal(savedMessage)
		if err != nil {
			log.Println("Error formatting message:", err)
			continue
		}

		// Broadcast the saved message
		manager.Broadcast(messageJSON)
	}
}
