package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	// Notifyで指定されたシグナルが来ると、チャンネルを通じて、そのシグナルを受け取れる。
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	s := <-signals
	// perf: goroutine
	switch s {
	case syscall.SIGINT:
		fmt.Println("SIGINT")
	case syscall.SIGTERM:
		fmt.Println("SIGTERM")
	}
}
