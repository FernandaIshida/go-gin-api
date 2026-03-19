package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panic("Error connecting to RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Panic("Error opening channel:", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"students_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic("Error declaring queue:", err)
	}

	msgs, err := ch.Consume(
		"students_queue",
		"",
		true, // auto ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic("Error consuming messages:", err)
	}

	fmt.Println("Waiting for messages...")

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			fmt.Println("Received message:", string(msg.Body))
		}
	}()

	<-forever
}
