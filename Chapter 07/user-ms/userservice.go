package main

import (
	"gomodules/configmodule"
	"log"
	"sync"
	"user-ms/server"
	"user-ms/server/dbserver"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting User Microservice")

	log.Println("Initializig configuration")
	err := configmodule.InitConfig("userservice", nil)
	if err != nil {
		log.Fatalf("Failed to read configuration: %v\n", err)
	}

	log.Println("Initializing database")
	err = dbserver.InitializeDb()
	if err != nil {
		log.Fatalf("Could not access database: %v\n", err)
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

		log.Println("Starting gRPC Server")

		err := server.StartGrpcServer()
		if err != nil {
			log.Fatalf("Could not start gRPC server: %v\n", err)
			return
		}

		log.Println("gRPC server gracefully terminated")
	}()

	wg.Wait()

	log.Println("User Microservice Stopped")
}
