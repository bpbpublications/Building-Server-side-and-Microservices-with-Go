package rabbitmq

import (
	"context"
	"os"
	"os/signal"
	"report-ms/config"
	"report-ms/data"
	"report-ms/server/dbserver"
	"syscall"
	"time"

	"github.com/streadway/amqp"
)

var (
	// StartRabbitMQConsumer starts rabbitMQ consumer
	StartRabbitMQConsumer = startRabbitMQConsumer
)

var channel *amqp.Channel
var queue amqp.Queue

func startRabbitMQConsumer() (err error) {
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

	consumeMessages()

	return
}

func consumeMessages() (err error) {
	ctx := context.Background()
	ctx = dbserver.PrepareDbRunner(ctx)

	// Position the first execution
	first := time.Now().Truncate(time.Hour * 24).Add(time.Hour * 17)
	if first.Before(time.Now()) {
		first = first.Add(time.Hour * 24)
	}
	firstC := time.After(first.Sub(time.Now()))

	// Receiving from a nil channel blocks forever
	ticker := &time.Ticker{C: nil}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-firstC:
		ticker = time.NewTicker(time.Hour * 24)
		err = getMessages(ctx)
		if err != nil {
			return
		}
	case <-ticker.C:
		err = getMessages(ctx)
		if err != nil {
			return
		}
	case <-interrupt:
		ticker.Stop()
		return
	}

	return
}

func getMessages(ctx context.Context) (err error) {
	messages, err := channel.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return
	}

	for message := range messages {
		err = data.CreateReport(ctx, string(message.Body))
		if err != nil {
			return
		}
	}

	return
}
