package main

import (
	"context"
	"github.com/nicholasjackson/env"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	l           *logrus.Logger
	bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")
)

func main() {

	// Parse the environment vars
	env.Parse()

	// Create logger(s)
	l = logrus.New()
	//stdLog := log.New(l.Writer(), "game-server", log.LstdFlags)

	m := CreateRouter()

	// Create a new server
	s := http.Server{
		Addr:         *bindAddress,      // configure the bind address
		Handler:      m,                 // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// Start the server
	go func() {
		l.Printf("Starting server on %s \n", *bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	l.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
