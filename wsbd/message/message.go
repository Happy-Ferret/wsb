package message

// Message struct websocket message
type Message struct {
	From string
	Body []byte
}

// NewMessage Return New Message
func NewMessage(from string, body []byte) Message {
	return Message{
		From: from,
		Body: body,
	}
}
