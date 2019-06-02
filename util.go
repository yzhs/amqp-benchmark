package main

import (
	"log"
	"strings"

	"github.com/streadway/amqp"
)

const (
	//ConnectionString string = "amqp://foobar:guest@eris:5672/"
	ConnectionString string = "amqp://guest:guest@localhost:5672/"
	NumberOfMessages int    = 10
)

var (
	MassTransitConnectionString string = strings.Replace(ConnectionString, "amqp://", "rabbitmq://", 1)
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ConnectToQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	durable := true
	deleteWhenUnused := false
	exclusive := false
	noWait := false

	q, err := ch.QueueDeclare(
		queueName,
		durable,
		deleteWhenUnused,
		exclusive,
		noWait,
		nil,
	)
	FailOnError(err, "Failed to declare a queue")

	return q
}
