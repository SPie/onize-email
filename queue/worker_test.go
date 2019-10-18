package queue

import (
    "errors"
    "sync"
    "testing"

    "github.com/spie/onize-email/email"
    "github.com/spie/onize-email/log"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestPullJobs(t *testing.T) {
    job1 := NewJob("Id1", "Name", "DisplayName", email.NewMessage("example1@email.com", "Content1"))
    job2 := NewJob("Id2", "Name", "DisplayName", email.NewMessage("example2@email.com", "Content2"))
    waitGroup := new(sync.WaitGroup)
    waitGroup.Add(2)
    deliveries := make(chan DeliveryContract, 2)
    defer close(deliveries)
    delivery1 := DeliveryMock{waitGroup: waitGroup}
    delivery1.On("GetJob").Return(job1, nil)
    delivery1.On("Ack").Return(nil)
    deliveries <- &delivery1
    delivery2 := DeliveryMock{waitGroup: waitGroup}
    delivery2.On("GetJob").Return(job2, nil)
    delivery2.On("Ack").Return(nil)
    deliveries <- &delivery2
    queueHandler := new(QueueHandlerMock)
    queueHandler.On("Consume").Return(deliveries, nil)
    emailHandler := new(EmailHandlerMock)
    emailHandler.On("SendEmail", "Name", job1.GetMessage()).Return(nil)
    emailHandler.On("SendEmail", "Name", job2.GetMessage()).Return(nil)
    log1 := log.NewLog(job1.Id, job1.Name, job1.Message.Recipient, job1.Message.Content)
    log2 := log.NewLog(job2.Id, job2.Name, job2.Message.Recipient, job2.Message.Content)
    logRepo := new(LogRepositoryMock)
    logRepo.On("Create", &log1).Return(logRepo)
    logRepo.On("Create", &log2).Return(logRepo)
    internalWaitGroup := new(sync.WaitGroup)
    internalWaitGroup.Add(1)

    go PullJobs(internalWaitGroup, queueHandler, emailHandler, logRepo)
    waitGroup.Wait()

    delivery1.AssertCalled(t, "Ack")
    delivery2.AssertCalled(t, "Ack")
    emailHandler.AssertCalled(t, "SendEmail", "Name", job1.GetMessage())
    emailHandler.AssertCalled(t, "SendEmail", "Name", job2.GetMessage())
    logRepo.AssertCalled(t, "Create", &log1)
    logRepo.AssertCalled(t, "Create", &log2)
}

func TestPullJobsWithErrorOnGetJob(t *testing.T) {
    waitGroup := new(sync.WaitGroup)
    waitGroup.Add(1)
    deliveries := make(chan DeliveryContract, 1)
    defer close(deliveries)
    delivery := DeliveryMock{waitGroup: waitGroup}
    delivery.On("GetJob").Return(Job{}, errors.New("GetJob error"))
    delivery.On("Reject").Return(nil)
    deliveries <- &delivery
    queueHandler := new(QueueHandlerMock)
    queueHandler.On("Consume").Return(deliveries, nil)
    emailHandler := new(EmailHandlerMock)
    internalWaitGroup := new(sync.WaitGroup)
    internalWaitGroup.Add(1)
    logRepo := new(LogRepositoryMock)

    go PullJobs(internalWaitGroup, queueHandler, emailHandler, logRepo)
    waitGroup.Wait()

    delivery.AssertNotCalled(t, "Ack")
    delivery.AssertCalled(t, "Reject")
    emailHandler.AssertNotCalled(t, "SendEmail")
}

func TestPullJobsWithErrorOnConsume(t *testing.T) {
    queueHandler := new(QueueHandlerMock)
    queueHandler.On("Consume").Return(make(chan DeliveryContract, 1), errors.New("Consume error"))
    emailHandler := new(EmailHandlerMock)
    logRepo := new(LogRepositoryMock)

    assert.Panics(t, func () {PullJobs(new(sync.WaitGroup), queueHandler, emailHandler, logRepo)})
}

func TestPullJobsWithErrorOnSendEmail(t *testing.T) {
    job := NewJob("Id1", "Name", "DisplayName", email.NewMessage("example1@email.com", "Content1"))
    waitGroup := new(sync.WaitGroup)
    waitGroup.Add(1)
    deliveries := make(chan DeliveryContract, 1)
    defer close(deliveries)
    delivery := DeliveryMock{waitGroup: waitGroup}
    delivery.On("GetJob").Return(job, nil)
    delivery.On("Reject").Return(nil)
    deliveries <- &delivery
    queueHandler := new(QueueHandlerMock)
    queueHandler.On("Consume").Return(deliveries, nil)
    emailHandler := new(EmailHandlerMock)
    emailHandler.On("SendEmail", "Name", job.GetMessage()).Return(errors.New("SendEmail error"))
    internalWaitGroup := new(sync.WaitGroup)
    internalWaitGroup.Add(1)
    logRepo := new(LogRepositoryMock)

    go PullJobs(internalWaitGroup, queueHandler, emailHandler, logRepo)
    waitGroup.Wait()

    delivery.AssertNotCalled(t, "Ack")
    delivery.AssertCalled(t, "Reject")
}

type DeliveryMock struct {
    mock.Mock
    waitGroup *sync.WaitGroup
}

func (delivery *DeliveryMock) Ack() error {
    defer delivery.waitGroup.Done()

    return delivery.Called().Error(0)
}

func (delivery *DeliveryMock) Reject() error {
    defer delivery.waitGroup.Done()

    return delivery.Called().Error(0)
}

func (delivery *DeliveryMock) GetJob() (Job, error) {
    args := delivery.Called()
    
    return args.Get(0).(Job), args.Error(1)
}

type QueueHandlerMock struct {
    mock.Mock
}

func (queueHandler QueueHandlerMock) Consume() (chan DeliveryContract, error) {
    args := queueHandler.Called()

    return args.Get(0).(chan DeliveryContract), args.Error(1)
}

func (queueHandler QueueHandlerMock) Close() error {
    return queueHandler.Called().Error(0)
}

type EmailHandlerMock struct {
    mock.Mock
}

func (emailHandler *EmailHandlerMock) SendEmail(identifier string, message email.Message) error {
    return emailHandler.Called(identifier, message).Error(0)
}

type LogRepositoryMock struct {
    mock.Mock
}

func (logRepo *LogRepositoryMock) Create(log *log.Log) log.LogRepository {
    logRepo.Called(log)

    return logRepo
}
