package main

import (
	"book-ms/server"
	"book-ms/server/dbserver"
	"book-ms/server/grpcserver"
	"book-ms/server/rabbitmq"
	"gomodules/configmodule"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting Book Microservice")

	log.Println("Initializig configuration")
	err := configmodule.InitConfig("bookservice", nil)
	if err != nil {
		log.Fatalf("Failed to read configuration: %v\n", err)
	}

	log.Println("Initializing database")
	err = dbserver.InitializeDb()
	if err != nil {
		log.Fatalf("Could not access database: %v\n", err)
	}

	log.Println("Starting User gRPC Client")
	err = grpcserver.StartUserGrpcClient()
	if err != nil {
		log.Fatalf("Could not start User gRPC Client: %v\n", err)
	}

	log.Println("Starting RabbitMQ Publisher")
	err = rabbitmq.StartRabbitMQPublisher()
	if err != nil {
		log.Fatalf("Could not start RabbitMQ Publisher: %v\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		log.Println("Starting HTTP Server")
		err := server.StartHTTPServer()
		if err != nil {
			log.Fatalf("Could not start HTTP Server: %v\n", err)
		}

		log.Println("HTTP Server gracefully terminated")
	}()

	wg.Wait()

	log.Println("Book Microservice Stopped")
}
