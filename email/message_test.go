package email

import (
    "testing"
    "encoding/json"

    "github.com/stretchr/testify/assert"
)

func TestCreateMessage(t *testing.T) {
    data := map[string]interface{}{
	"key1": "value",
	"key2": map[string]string{"subKey1": "subValue1"},
    }
    message := NewMessage("Recipient", data)
    assert.Equal(t, "Recipient", message.GetRecipient())
    assert.Equal(t, data, message.GetData())
}

func TestParseMessageFromJson(t *testing.T) {
    jsonString := `{
	"recipient": "Recipient",
	"data": {
	    "key1": "value",
	    "key2": {
		"subKey1": "subValue1"
	    }
	}
    }`
    var message Message
    err := json.Unmarshal([]byte(jsonString), &message)
    assert.Empty(t, err)
    assert.Equal(t, "Recipient", message.GetRecipient())
    assert.Equal(
	t,
	map[string]interface{}{
	    "key1": "value",
	    "key2": map[string]interface{}{"subKey1": "subValue1"},
        },
	message.GetData(),
    )
}
