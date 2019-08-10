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
    sender := "" //TODO
    emailUsername := "" //TODO
    emailPassword := "" //TODO
    emailHost := "" //TODO
    emailPort := "" //TODO
    templatesDir := "templates"

    waitGroup := new(sync.WaitGroup)

    queueHandler, err := queue.OpenQueue(host, port, username, password, queueName)
    failOnError(err, "")
    defer queueHandler.Close()

    emailHandler := email.NewEmailHandler(
	sender,
	emailHost,
	emailPort,
	email.NewAuthUser(emailUsername, emailPassword, emailHost),
	email.NewParser(templatesDir),
    )

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
