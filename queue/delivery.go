package queue

import (
    "encoding/json"

    "github.com/streadway/amqp"
)

type DeliveryContract interface {
    Ack() error
    Reject() error
    GetJob() (Job, error)
}

type Delivery struct {
    amqpDelivery amqp.Delivery
}

func NewDelivery(amqpDelivery amqp.Delivery) Delivery {
    return Delivery{amqpDelivery: amqpDelivery}
}

func (delivery Delivery) Ack() error {
    return delivery.amqpDelivery.Ack(false)
}

func (delivery Delivery) Reject() error {
    return delivery.amqpDelivery.Reject(true)
}

func (delivery Delivery) GetJob() (Job, error) {
    var job Job
    err := json.Unmarshal(delivery.amqpDelivery.Body, &job)
    if err != nil {
	return Job{}, err
    }

    return job, nil
}
