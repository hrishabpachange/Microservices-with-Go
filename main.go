package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hrishabpachange/go-basic-API/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l) //instance of products.go

	sm := http.NewServeMux() // Custom Serve Multiplexer (ServeMux)
	sm.Handle("/", ph)

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		fmt.Println("Server running...")

		err := s.ListenAndServe()

		if err != nil { //Error and Exception handling in Go
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)      //This creates channel that can receive OS signals
	signal.Notify(sigChan, os.Interrupt)    //Whenever an OS interrupt happens, send it to sigchan(os.interrupt is triggered when ctrl+c is pressed)
	signal.Notify(sigChan, syscall.SIGTERM) //Notify when program receives SIGTERM signal

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second) // we are giving the server 30 seconds of time to close all the handlers after which if any handlers are running it will close them forcefully
	defer cancel()                                                          // ensure the cancel function is called to avoid context leak
	s.Shutdown(tc)                                                          //Graceful shutdown
}
