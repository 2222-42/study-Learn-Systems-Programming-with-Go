package main

import (
	"context"
	"fmt"
	"github.com/lestrrat/go-server-starter/listener"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	listeners, err := listener.ListenAll()
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "server pid: %d %v\n", os.Getpid(), os.Environ())
		}),
	}
	go server.Serve(listeners[0])

	<-signals
	server.Shutdown(context.Background())
}
