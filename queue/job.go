package queue

import (
    "github.com/spie/onize-email/email"
)

type Job struct {
    Id string `json:"id"`
    Name string `json:"job"`
    DisplayName string `json:"displayName"`
    Message email.Message `json:"data"`
}

func NewJob(id string, name string, displayName string, message email.Message) Job {
    return Job{Id: id, Name: name, DisplayName: displayName, Message: message}
}

func (job Job) GetId() string {
    return job.Id
}

func (job Job) GetName() string {
    return job.Name
}

func (job Job) GetDisplayName() string {
    return job.DisplayName
}

func (job Job) GetMessage() email.Message {
    return job.Message
}
