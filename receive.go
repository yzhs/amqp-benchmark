package main

import (
	"log"
	"os"

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

	GetMessages(msgs, 100000, queue.Name)
}

func GetMessages(msgs <-chan amqp.Delivery, maxNumber int, queueName string) {
	f, err := os.Create("queues/" + queueName)
	FailOnError(err, "Could not create file")
	defer f.Close()

	count := 0
	for msg := range msgs {
		_, err = f.Write(msg.Body)
		FailOnError(err, "Could not write to file")
		_, err = f.Write([]byte("\n\n"))
		FailOnError(err, "Could not write to file")

		count += 1
		if count >= maxNumber {
			break
		}
	}

	log.Printf("Received %d messages", count)
}
