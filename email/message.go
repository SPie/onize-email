package email

type Message struct {
    Recipient string `json:"recipient"`
    Content string `json:"content"`
}

func NewMessage(recipient string, content string) Message {
    return Message{
	Recipient: recipient,
	Content: content,
    }   
}

func (message Message) GetRecipient() string {
    return message.Recipient
}

func (message Message) GetContent() string {
    return message.Content
}
