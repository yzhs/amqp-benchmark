package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://foobar:guest@eris:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	durable := true
	deleteWhenUnused := false
	exclusive := false
	noWait := false
	q, err := ch.QueueDeclare(
		"pingpong",
		durable,
		deleteWhenUnused,
		exclusive,
		noWait,
		nil,
	)
	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	count := 0
	for _ = range msgs {
		count += 1
		if count >= 100000 {
			break
		}
	}

	log.Printf("Received %d messages", count)
}
