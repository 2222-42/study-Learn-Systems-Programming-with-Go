package main

import (
	"fmt"
	"time"
)

// 4.1
func main() {
	fmt.Println("start sub()")
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()

	time.Sleep(2 * time.Second)
}
