package queue

import (
    "testing"
    "encoding/json"

    "github.com/spie/onize-email/email"
    "github.com/stretchr/testify/assert"
)

func TestCreateJob(t *testing.T) {
    message := email.Message{Recipient: "Recipient", Content: "Content"}
    job := NewJob("Id", "JobName", "Display name", message)
    assert.Equal(t, "Id", job.GetId())
    assert.Equal(t, "JobName", job.GetName())
    assert.Equal(t, "Display name", job.GetDisplayName())
    assert.Equal(t, message, job.GetMessage())
}

func TestParseJobFromJson(t *testing.T) {
    jsonString := `{
	"id": "Id",
	"job": "JobName",
	"displayName": "Display name",
	"data": {
	    "recipient": "Recipient",
	    "content" : "Content"
	}
    }`
    var job Job
    err := json.Unmarshal([]byte(jsonString), &job)
    assert.Empty(t, err)
    assert.Equal(t, "Id", job.GetId())
    assert.Equal(t, "JobName", job.GetName())
    assert.Equal(t, "Display name", job.GetDisplayName())
    assert.Equal(t, email.NewMessage("Recipient", "Content"), job.GetMessage())
}
