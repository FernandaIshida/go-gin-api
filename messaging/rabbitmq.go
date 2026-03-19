package messaging

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panic("Error connecting to RabbitMQ:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Panic("Error opening channel:", err)
	}

	return conn, ch
}

func DeclareQueue(ch *amqp.Channel, queueName string) {
	_, err := ch.QueueDeclare(
		queueName,
		true,  // durable (restart survive)
		false, // auto delete
		false, // exclusive
		false, // no wait
		nil,   // args
	)

	if err != nil {
		log.Panic("Error declaring queue:", err)
	}
}

func PublishMessage(queueName string, body string) {
	conn, ch := ConnectRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	DeclareQueue(ch, queueName)

	err := ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		log.Panic("Error publishing message:", err)
	}
}
