package main

import (
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

	body := []byte("ping")
	for i := 0; i < NumberOfMessages; i++ {
		send(ch, queue, body)
	}

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
