package main

import (
	"fmt"
	"time"
)

// 4.5
// Q4.1
func main() {
	fmt.Println("Start.")

	<-time.After(10 * time.Second)

	fmt.Println("It pass 10 seconds.")
}
