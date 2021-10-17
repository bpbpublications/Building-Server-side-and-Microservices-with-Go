package main

import (
	"building-restful-web-services-with-go/chapter2/library/config"
	"building-restful-web-services-with-go/chapter2/library/server"
	"building-restful-web-services-with-go/chapter2/library/server/dbserver"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting Library Server")

	log.Println("Initializig configuration")
	err := config.InitConfig("library", nil)
	if err != nil {
		log.Fatalf("Failed to read configuration: %v\n", err)
	}

	log.Println("Initializing database")
	err = dbserver.InitializeDb()
	if err != nil {
		log.Fatalf("Could not access database: %v\n", err)
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

	log.Println("Library Server Stopped")
}
