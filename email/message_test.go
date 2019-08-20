package email

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateMessage(t *testing.T) {
    message := NewMessage("Recipient", "Content")
    assert.Equal(t, "Recipient", message.GetRecipient())
    assert.Equal(t, "Content", message.GetContent())
}

func TestParseMessageFromJson(t *testing.T) {
    jsonString := `{
	"recipient": "Recipient",
	"content": "Content"
    }`
    var message Message
    err := json.Unmarshal([]byte(jsonString), &message)
    assert.Empty(t, err)
    assert.Equal(t, "Recipient", message.GetRecipient())
    assert.Equal( t, "Content", message.GetContent())
}
