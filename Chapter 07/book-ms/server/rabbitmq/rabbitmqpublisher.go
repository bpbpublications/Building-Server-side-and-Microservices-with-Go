package rabbitmq

import (
	"book-ms/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/streadway/amqp"
)

var (
	// StartRabbitMQPublisher starts rabbitMQ publisher
	StartRabbitMQPublisher = startRabbitMQPublisher

	// SendMessage will send message to queue
	SendMessage = sendMessage
)

var channel *amqp.Channel
var queue amqp.Queue

func startRabbitMQPublisher() (err error) {
	address := config.GetRabbitMQConnectionString()

	conn, err := amqp.Dial(address)
	if err != nil {
		return
	}

	channel, err = conn.Channel()
	if err != nil {
		return
	}

	queue, err = channel.QueueDeclare("books", false, false, false, false, nil)
	if err != nil {
		return
	}

	go func() {
		// Listen to operating system's interrupt signal
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		<-interrupt

		// Gracefully shut down the publisher when it happens
		conn.Close()
		channel.Close()
	}()

	return
}

func sendMessage(bookID string) (err error) {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(bookID),
	}

	err = channel.Publish("", queue.Name, false, false, message)
	if err != nil {
		return
	}

	return
}
