package queue

import (
    "sync"

    "github.com/spie/onize-email/email"
)

func PullJobs(waitGroup *sync.WaitGroup, queueHandler QueueHandlerContract, emailHandler email.EmailHandlerContract) {
    deliveries, err := queueHandler.Consume()
    if err != nil {
	panic("queueHandler.Consume() failed")
    }

    for {
	delivery, open := <-deliveries
	if !open {
	    waitGroup.Done()
	    return
	}

	job, err := delivery.GetJob()
	if err != nil {
	    delivery.Reject(true)
	    continue
	}

	emailHandler.SendEmail(job.GetId(), job.GetMessage())

	delivery.Ack(false)
    }
}
