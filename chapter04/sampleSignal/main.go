package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 4.3.1
func main() {
	signals := make(chan os.Signal, 1)
	//go func() {
	signal.Notify(signals, syscall.SIGINT)
	//}() // OSから受け取ったシグナルを受け取るまで待つ

	fmt.Println("Waiting SIGINT (CTRL+C)")
	<-signals
	fmt.Println("SIGINT arrived")
}
