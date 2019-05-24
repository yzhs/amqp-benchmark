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

	queue := ConnectToQueue(ch, "pingpong")

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	FailOnError(err, "Failed to register a consumer")

	GetMessages(msgs, 100000)
}

func GetMessages(msgs <-chan amqp.Delivery, maxNumber int) {
	count := 0
	for _ = range msgs {
		count += 1
		if count >= maxNumber {
			break
		}
	}

	log.Printf("Received %d messages", count)
}
