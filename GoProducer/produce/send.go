package produce

import (
	"log"

	"github.com/streadway/amqp"
)

var message string

func PostMessage(msg string) {
	message = msg
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	connectToRabbitMQServer(conn)
}

func connectToRabbitMQServer(connection *amqp.Connection) {
	ch, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	sendMessage(ch)
}

func sendMessage(channel *amqp.Channel) {
	q, err := channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := message
	err = channel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
