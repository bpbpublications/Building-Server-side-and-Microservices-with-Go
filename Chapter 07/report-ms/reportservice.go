package main

import (
	"log"
	"report-ms/server"
	"report-ms/server/dbserver"
	"report-ms/server/grpcserver"
	"report-ms/server/rabbitmq"
	"sync"

	"gomodules/configmodule"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting Report Microservice")

	log.Println("Initializig configuration")
	err := configmodule.InitConfig("reportservice", nil)
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

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		log.Println("Starting HTTP Server")
		err := server.StartHTTPServer()
		if err != nil {
			log.Fatalf("Could not start HTTP Server: %v\n", err)
		}

		log.Println("HTTP Server gracefully terminated")
	}()

	go func() {
		defer wg.Done()

		log.Println("Starting RabbitMQ Consumer")
		err = rabbitmq.StartRabbitMQConsumer()
		if err != nil {
			log.Fatalf("Could not start RabbitMQ Consumer: %v\n", err)
		}

		log.Println("RabbitMQ Consumer gracefully terminated")
	}()

	wg.Wait()

	log.Println("Report Microservice Stopped")
}
