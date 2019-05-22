package main

import (
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

	body := []byte("ping")
	for i := 0; i < 100000; i++ {
		send(ch, q, body)
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
