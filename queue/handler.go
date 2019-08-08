package queue

import (
    "fmt"

    "github.com/streadway/amqp"
)

type QueueHandlerContract interface {
    Consume() (chan DeliveryContract, error)
    Close() error
}

type QueueHandler struct {
    connection *amqp.Connection
    channel *amqp.Channel
    queue amqp.Queue
}

func OpenQueue(host string, port string, username string, password string, queueName string) (QueueHandlerContract, error) {
    url := fmt.Sprintf("amqp://%s:%s@%s:%s", username, password, host, port)
    connection, err := amqp.Dial(url)
    if err != nil {
	return nil, err
    }
    
    channel, err := connection.Channel()
    if err != nil {
	return nil, err
    }

    queue, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
    if err != nil {
	return nil, err
    }

    err = channel.Qos(1, 0, false)
    if err != nil {
	return nil, err
    }

    return QueueHandler{connection: connection, channel: channel, queue: queue}, nil
}

func (queueHandler QueueHandler) Consume() (chan DeliveryContract, error) {
    amqpDeliveries, err := queueHandler.channel.Consume(queueHandler.queue.Name, "", false, false, false, false, nil)
    if err != nil {
	return nil, err
    }

    deliveries := make(chan DeliveryContract)

    go func () {
	for amqpDelivery := range amqpDeliveries {
	    deliveries <- NewDelivery(amqpDelivery)
        }
    }()

    return deliveries, nil
}

func (queueHandler QueueHandler) Close() error {
    defer queueHandler.connection.Close()
    defer queueHandler.channel.Close()

    return nil
}

func failOnError(err error, errorMessage string) {
    if err != nil {
	panic(errorMessage)
    }
}
