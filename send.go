package main

import (
	"strconv"

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

	for i := 0; i < NumberOfMessages; i++ {
		body := []byte(generateBody(i))
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
			ContentType: "application/json",
			Body:        body,
		})
	FailOnError(err, "Failed to publish a message")
}

func generateBody(i int) string {
	return `{
	"destinationAddress": "` + MassTransitConnectionString + `",
	"headers": {},
	"message": {
		"value": ` + strconv.Itoa(i) + `
	},
	"messageType"
}`
}
