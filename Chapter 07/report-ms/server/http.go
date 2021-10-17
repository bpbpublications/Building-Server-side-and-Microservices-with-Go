package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"report-ms/config"
	"syscall"
)

var (
	// StartHTTPServer starts listening to HTTP requests
	StartHTTPServer = startHTTPServer
)

func startHTTPServer() (err error) {
	mux := http.NewServeMux()
	mux.Handle("/api/", newHandlerAPI())

	server := http.Server{}

	server.ReadTimeout = config.GetHTTPReadTimeout()
	server.WriteTimeout = config.GetHTTPWriteTimeout()
	server.Addr = config.GetHTTPServerAddress()
	server.Handler = mux

	go func() {
		// Listen to operating system's interrupt signal
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		<-interrupt

		// Gracefully shut down the server when it happens
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down. %v\n", err)
		}
	}()

	err = server.ListenAndServe()
	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
