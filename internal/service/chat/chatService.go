package chat

import (
	"encoding/json"
	"pumplepet-server/internal/model"
	"pumplepet-server/pkg/database"
)

// MessageBroadcaster interface to avoid circular imports
type MessageBroadcaster interface {
	Broadcast(message []byte)
}

func GetMessages() ([]model.Message, error) {
	var messages []model.Message
	if err := database.DB.Preload("User").Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func SaveMessage(content string, senderID uint, recipientID uint) (*model.Message, error) {
	message := &model.Message{
		Content:     content,
		SenderID:    senderID,
		RecipientID: recipientID,
	}

	if err := database.DB.Create(message).Error; err != nil {
		return nil, err
	}

	// Load both sender and recipient data
	if err := database.DB.Preload("Sender").Preload("Recipient").First(message, message.ID).Error; err != nil {
		return nil, err
	}

	return message, nil
}

func BroadcastMessage(broadcaster MessageBroadcaster, message *model.Message) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	broadcaster.Broadcast(messageJSON)
	return nil
}

func GetMessagesBetweenUsers(userID1, userID2 uint) ([]model.Message, error) {
    var messages []model.Message
    err := database.DB.
        Preload("Sender").
        Preload("Recipient").
        Where(
            "(sender_id = ? AND recipient_id = ?) OR (sender_id = ? AND recipient_id = ?)",
            userID1, userID2, userID2, userID1,
        ).
        Order("created_at asc").
        Find(&messages).Error
    return messages, err
}