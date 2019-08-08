package email

type Message struct {
    Recipient string `json:"recipient"`
    Data map[string]interface{} `json:"data"`
}

func NewMessage(recipient string, data map[string]interface{}) Message {
    return Message{
	Recipient: recipient,
	Data: data,
    }   
}

func (message Message) GetRecipient() string {
    return message.Recipient
}

func (message Message) GetData() map[string]interface{} {
    return message.Data
}
