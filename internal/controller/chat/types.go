package chat

type MessagePayload struct {
	Content     string `json:"content"`
	RecipientID uint   `json:"recipient_id"`
}
