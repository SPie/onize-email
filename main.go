package main

import (
    "sync"

    "github.com/spie/onize-email/email"
    "github.com/spie/onize-email/queue"
)

func main() {
    workersCount := 3
    host := "127.0.0.1"
    port := "5672"
    username := "guest"
    password := "guest"
    queueName := "email"

    waitGroup := new(sync.WaitGroup)

    queueHandler, err := queue.OpenQueue(host, port, username, password, queueName)
    failOnError(err, "")
    defer queueHandler.Close()

    emailHandler := email.NewEmailHandler()

    for i := 1; i <= workersCount; i++ {
	waitGroup.Add(1)

	go queue.PullJobs(waitGroup, queueHandler, emailHandler)
    }

    waitGroup.Wait()
}

func failOnError(err error, errorMessage string) {
    if err == nil {
	return
    }

    if errorMessage == "" {
	errorMessage = err.Error()
    }

    panic(errorMessage)
}
