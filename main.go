package main

import (
    "os"
    "strconv"
    "sync"

    "github.com/joho/godotenv"
    "github.com/spie/onize-email/email"
    "github.com/spie/onize-email/queue"
)

func main() {
    loadEnv()

    waitGroup := new(sync.WaitGroup)

    queueHandler, err := queue.OpenQueue(os.Getenv("QUEUE_HOST"), os.Getenv("QUEUE_PORT"), os.Getenv("QUEUE_USERNAME"), os.Getenv("QUEUE_PASSWORD"), os.Getenv("QUEUE_NAME"))
    failOnError(err, "")
    defer queueHandler.Close()

    failOnError(err, "")

    emailHandler := email.NewEmailHandler(
	os.Getenv("EMAIL_SENDER"),
	os.Getenv("EMAIL_HOST"),
	os.Getenv("EMAIL_PORT"),
	email.NewAuthUser(os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("EMAIL_HOST")),
    )

    workersCount, err := strconv.Atoi(os.Getenv("WORKERS_COUNT"))
    if err != nil {
	failOnError(err, "")
    }

    for i := 1; i <= workersCount; i++ {
	waitGroup.Add(1)

	go queue.PullJobs(waitGroup, queueHandler, emailHandler)
    }

    waitGroup.Wait()
}

func loadEnv() {
    err := godotenv.Load()
    failOnError(err, "")
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
