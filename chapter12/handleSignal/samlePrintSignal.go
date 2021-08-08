package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	signals := make(chan os.Signal, 1)
	// Notifyで指定されたシグナルが来ると、チャンネルを通じて、そのシグナルを受け取れる。
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-signals
		// perf: goroutine
		switch s {
		case syscall.SIGINT:
			fmt.Println("SIGINT")
			defer wg.Done()
		case syscall.SIGTERM:
			fmt.Println("SIGTERM")
			defer wg.Done()
		}
	}()

	fmt.Println("print signal for 5 seconds.")
	time.Sleep(5 * time.Second)

	//fmt.Println("no printing now.")
	//signal.Reset(syscall.SIGINT, syscall.SIGHUP)

	fmt.Println("back to default")
	signal.Stop(signals)

	wg.Wait()
}
