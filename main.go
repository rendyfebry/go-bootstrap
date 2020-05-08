package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rendyfebry/go-streamer/transport"
)

func main() {
	env := "dev"
	host := "localhost"
	port := 3000

	routes := transport.MakeHTTPRoutes()

	// Initialize http serve
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routes,
	}

	log.Println("Starting!")
	log.Printf("- Environment %s \n", env)
	log.Printf("- Application URL http://%s:%d \n", host, port)

	// Run server in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
