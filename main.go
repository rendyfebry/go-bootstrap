package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rendyfebry/go-streamer/service"
	"github.com/rendyfebry/go-streamer/transport"
	"github.com/rendyfebry/go-streamer/utils"
)

func main() {
	// Init config
	cfg := utils.GetConfig()

	// Init service
	svc := service.NewService(cfg)

	// Init http transport route
	routes := transport.MakeHTTPRoutes(svc)

	// Init http server
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routes,
	}

	log.Println("Starting!")
	log.Printf("- Environment %s \n", cfg.Env)
	log.Printf("- Application URL http://%s:%d \n", cfg.Host, cfg.Port)

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
