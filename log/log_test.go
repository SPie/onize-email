package log

import (
    "testing"

    "github.com/spie/onize-email/db"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestNewLog(t *testing.T) {
    log := NewLog("JobId", "JobName", "Recipient", "Content")

    assert.Equal(t, "JobId", log.JobId)
    assert.Equal(t, "JobName", log.JobName)
    assert.Equal(t, "Recipient", log.Recipient)
    assert.Equal(t, "Content", log.Content)
}

func TestLogRepositoryCreate(t *testing.T) {
    log := &Log{}
    connection := &TestConnection{}
    connection.On("Create", log).Return(connection)
    logRepo := &GormLogRepository{connection: connection}

    logRepo.Create(log)

    connection.AssertCalled(t, "Create", log)
}

type TestConnection struct {
    mock.Mock
}

func (connection *TestConnection) Close() error {
    return nil
}

func (connection *TestConnection) Create(model interface{}) db.Connection {
    connection.Called(model)
    return connection
}
