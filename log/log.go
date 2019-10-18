package log

import (
    "time"

    "github.com/spie/onize-email/db"
)

type Log struct {
    ID int
    JobId string
    JobName string
    Recipient string
    Content string
    CreatedAt *time.Time
}

func NewLog(jobId string, jobName string, recipient string, content string) Log {
    return Log{
	JobId: jobId,
	JobName: jobName,
	Recipient: recipient,
	Content: content,
    }
}

type LogRepository interface{
    Create(log *Log) LogRepository
}

type GormLogRepository struct {
    connection db.Connection    
}

func (logRepo *GormLogRepository) Create(log *Log) LogRepository {
    logRepo.connection.Create(log)

    return logRepo
}

func NewRepository(connection db.Connection) *GormLogRepository {
    return &GormLogRepository{connection: connection}
}


