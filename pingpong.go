package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial(ConnectionString)
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

	count := 0
	body := []byte("ping")

	for i := 0; i < NumberOfMessages/2; i++ {
		send(ch, queue, body)
		_ = <-msgs
		count += 1
	}

	log.Printf("Received %d messages", count)
}

func send(ch *amqp.Channel, q amqp.Queue, body []byte) {
	err := ch.Publish(
		"",     //exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	FailOnError(err, "Failed to publish a message")
}
